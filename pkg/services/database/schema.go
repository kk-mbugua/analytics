package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log" // Import log package
)

// TableInfo represents the structure of table and column information
type TableInfo struct {
	TableName  string      `json:"table_name"`
	Columns    []ColumnInfo `json:"columns"` // Updated to use ColumnInfo
}

// ColumnInfo represents the structure of column information
type ColumnInfo struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}

// SchemaInfo represents the structure of schema information including tables
type SchemaInfo struct {
	SchemaName string      `json:"schema_name"`
	Tables     []TableInfo `json:"tables"`
}

// InitializeDB initializes the database connection
func InitializeDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=p3rs0n4l dbname=datest1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err) // Changed fmt.Printf to log.Printf
		return nil
	}
	return db
}

// GetSchemas retrieves the schemas from the database
func GetSchemas(db *gorm.DB) ([]SchemaInfo, error) {
	var schemas []SchemaInfo
	querySchemas := "SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('information_schema', 'pg_catalog', 'pg_toast')"

	var schemaNames []struct{ SchemaName string }
	result := db.Raw(querySchemas).Scan(&schemaNames)
	if result.Error != nil {
		log.Printf("Error retrieving schema names: %v", result.Error) // Log schema retrieval error
		return nil, result.Error
	}

	for _, schema := range schemaNames {
		var tableInfo []TableInfo
		queryTables := fmt.Sprintf(`SELECT table_name, column_name, data_type FROM information_schema.columns WHERE table_schema = '%s' ORDER BY table_name, ordinal_position;`, schema.SchemaName)

		var columns []struct {
			TableName  string
			ColumnName string
			DataType   string
		}
		result = db.Raw(queryTables).Scan(&columns)
		if result.Error != nil {
			log.Printf("Error retrieving tables for schema (%s): %v", schema.SchemaName, result.Error) // Log error for each schema
			return nil, result.Error
		}

		// Grouping columns by table name
		tableMap := make(map[string][]ColumnInfo)
		for _, col := range columns {
			tableMap[col.TableName] = append(tableMap[col.TableName], ColumnInfo{
				ColumnName: col.ColumnName,
				DataType:   col.DataType,
			})
		}

		// Create SchemaInfo with populated Tables and Columns
		for tableName, cols := range tableMap {
			tableInfo = append(tableInfo, TableInfo{
				TableName: tableName,
				Columns:   cols,
			})
		}

		schemas = append(schemas, SchemaInfo{SchemaName: schema.SchemaName, Tables: tableInfo})
	}
	return schemas, nil
}

// Database struct to hold the database connection logic
type Database struct{}

// Connect method to establish a connection to the PostgreSQL database
func (d *Database) Connect(dbURL string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database with URL %s: %v", dbURL, err) // Log error on connection failure
		return nil, err
	}
	return gormDB, nil
}
