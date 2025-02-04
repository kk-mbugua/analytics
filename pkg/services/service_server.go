package services

import (
    "context"
	database "analytics/pkg/services/database"
    pb "analytics/pkg/pb"
    "encoding/csv"
    "fmt"
    "log"
    "strings"
    "io"
    "bytes"
)

// SchemaServiceServer struct with the database dependency
type SchemaServiceServer struct {
    pb.UnimplementedSchemaServiceServer
    Database *database.Database
}

// GetSchemas retrieves schemas from the database
func (s *SchemaServiceServer) GetSchemas(ctx context.Context, req *pb.SchemaRequest) (*pb.SchemaResponse, error) {
    var allSchemas []*pb.SchemaInfo

    for _, dbURL := range req.DbUrls {
        // Attempt to connect to the database
        gormDB, err := s.Database.Connect(dbURL)
        if err != nil {
            log.Printf("Error connecting to database (%s): %v", dbURL, err)
            return nil, err
        }
		log.Printf("Successfully connected to database: %s", dbURL)

        // Retrieve schemas
        schemaInfos, err := database.GetSchemas(gormDB)
        if err != nil {
            log.Printf("Error retrieving schemas for database (%s): %v", dbURL, err)
            return nil, err
        }

        // Convert each schema from database.SchemaInfo to pb.SchemaInfo
        for _, schemaInfo := range schemaInfos {
            pbSchema := &pb.SchemaInfo{
                SchemaName: schemaInfo.SchemaName,
                Tables:     convertTables(schemaInfo.Tables),
            }
            allSchemas = append(allSchemas, pbSchema)
        }
    }

    response := &pb.SchemaResponse{Schemas: allSchemas}
    return response, nil
}

// GetTableData retrieves rows and columns from a specified table using CSV filters
func (s *SchemaServiceServer) GetTableData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
    var allTableData []*pb.TableData

    for _, dbURL := range req.DbUrls {
        gormDB, err := s.Database.Connect(dbURL)
        if err != nil {
            log.Printf("Error connecting to database (%s): %v", dbURL, err)
            return nil, err
        }

        for _, tableName := range req.TableNames {
            var rows []map[string]interface{}
            
            // Build the query by selecting specific columns
            query := gormDB.Table(tableName).Select(strings.Join(req.Columns, ", "))

            // Apply limit and offset
            query = query.Limit(int(req.Limit)).Offset(int(req.Offset))

            // Execute the query
            if err := query.Find(&rows).Error; err != nil {
                log.Printf("Error retrieving data from table (%s): %v", tableName, err)
                return nil, err
            }

            // Convert map[string]interface{} to pb.RowData
            var rowDataList []*pb.RowData
            for _, row := range rows {
                rowData := &pb.RowData{Columns: make(map[string]string)}
                for columnName, value := range row {
                    rowData.Columns[columnName] = fmt.Sprintf("%v", value) // Convert value to string
                }
                rowDataList = append(rowDataList, rowData)
            }

            allTableData = append(allTableData, &pb.TableData{
                TableName: tableName,
                Rows:      rowDataList,
            })
        }
    }

    response := &pb.DataResponse{TableData: allTableData}
    return response, nil
}


// Helper function to convert []database.TableInfo to []*pb.TableInfo
func convertTables(tables []database.TableInfo) []*pb.TableInfo {
    var pbTables []*pb.TableInfo
    for _, table := range tables {
        pbTables = append(pbTables, &pb.TableInfo{
            TableName: table.TableName,
            Columns:   convertColumns(table.Columns), // Convert the Columns properly
        })
    }
    return pbTables
}

// Helper function to convert []database.ColumnInfo to []*pb.ColumnInfo
func convertColumns(columns []database.ColumnInfo) []*pb.ColumnInfo {
    var pbColumns []*pb.ColumnInfo
    for _, column := range columns {
        pbColumns = append(pbColumns, &pb.ColumnInfo{
            ColumnName: column.ColumnName,
            DataType:   column.DataType,
        })
    }
    return pbColumns
}

// extractColumnsFromCSV parses CSV data and returns the column names
func extractColumnsFromCSV(csvData []byte) ([]string, error) {
    reader := csv.NewReader(bytes.NewReader(csvData))

    // Read the first line of the CSV (header)
    header, err := reader.Read()
    if err != nil {
        return nil, err
    }

    return header, nil
}

// UploadCsv handles CSV file uploads via gRPC streaming
func (s *SchemaServiceServer) UploadCsv(stream pb.SchemaService_UploadCsvServer) error {
    var buffer bytes.Buffer

    // Read the incoming stream
    for {
        chunk, err := stream.Recv()
        if err == io.EOF {
            // When done receiving, process the CSV data
            columns, err := extractColumnsFromCSV(buffer.Bytes())
            if err != nil {
                return err
            }

            // Parse the rows data from CSV file
            rows, err := extractRowsFromCSV(buffer.Bytes(), columns)
            if err != nil {
                return err
            }

            // Return the columns and rows as the response
            return stream.SendAndClose(&pb.CsvResponse{
                Columns: columns, // Return columns
                Rows:    rows,    // Return rows
            })
        }
        if err != nil {
            return err
        }

        // Write chunk data to the buffer
        buffer.Write(chunk.ChunkData)
    }
}

// extractRowsFromCSV parses the CSV data and returns the rows
func extractRowsFromCSV(csvData []byte, columns []string) ([]*pb.RowData, error) {
    reader := csv.NewReader(bytes.NewReader(csvData))

    // Read all rows of the CSV data
    var rows []*pb.RowData
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }

        // Map each row's column value to a corresponding Column name
        rowData := &pb.RowData{Columns: make(map[string]string)}
        for i, value := range record {
            if i < len(columns) {
                rowData.Columns[columns[i]] = value
            }
        }
        rows = append(rows, rowData)
    }

    return rows, nil
}

// GetFilteredCsvData filters CSV data based on requested columns and row limits
func (s *SchemaServiceServer) GetFilteredCsvData(ctx context.Context, req *pb.CsvRequest) (*pb.CsvResponse, error) {
	reader := csv.NewReader(bytes.NewReader(req.ChunkData))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSV data: %w", err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	// Extract header
	header := rows[0]
	columnIndices := make(map[string]int)
	for i, col := range header {
		columnIndices[col] = i
	}

	// Filter columns
	filteredHeader := []string{}
	for _, col := range req.Columns {
		if _, exists := columnIndices[col]; exists {
			filteredHeader = append(filteredHeader, col)
		} else {
			return nil, fmt.Errorf("column %s not found in CSV", col)
		}
	}

	// Filter rows
	filteredRows := []*pb.RowData{}
	for i, row := range rows[1:] { // Skip header
		if req.RowLimit > 0 && i >= int(req.RowLimit) {
			break
		}
		filteredRow := &pb.RowData{Columns: make(map[string]string)}
		for _, col := range req.Columns {
			filteredRow.Columns[col] = row[columnIndices[col]]
		}
		filteredRows = append(filteredRows, filteredRow)
	}

	return &pb.CsvResponse{
		Columns: filteredHeader,
		Rows:    filteredRows,
	}, nil
}
