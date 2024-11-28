package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	pb "analytics/pkg/pb"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSchemaServiceClient(conn)

	// Load a test CSV file
	csvFile := "/Users/work/dev/cigona/dail_afrika/analytics/da-ms-template/pkg/services/client_test/csv/test.csv"
	csvData, err := os.ReadFile(csvFile)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Prepare the request for filtering
	req := &pb.CsvRequest{
		ChunkData: csvData,
		Columns:   []string{"sorta", "thing"}, // Specify desired columns
		RowLimit:  5,                             // Limit rows
	}

	// Call GetFilteredCsvData
	resp, err := client.GetFilteredCsvData(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetFilteredCsvData: %v", err)
	}

	// Print response
	fmt.Println("Filtered Columns:", resp.Columns)
	for i, row := range resp.Rows {
		fmt.Printf("Row %d: %v\n", i+1, row.Columns)
	}
}
