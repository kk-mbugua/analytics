package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func HandleMigrateError(modelName string, err error) {
	if err != nil {
		log.Fatalf("Failed to auto migrate %s database: %v", modelName, err)
	}
	fmt.Printf("%s Database connected successfully...\n", modelName)
}

func InitDB(host string, port string, dbName string, dbUser string, password string) {
	// Open DB connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed %v", err)
	}
	// Migrate the schema
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Debug()
	if false {
		// Handle Schema Migrations Here for each model eg. handleMigrateError("Pipeline", tx.AutoMigrate(&pipelines.Pipeline{}))
		log.Printf("Migrating Pipeline")
	}
	tx.Commit()
	fmt.Print("Database connected successfully...")

	fmt.Println("DB Connection Successful...")

}
