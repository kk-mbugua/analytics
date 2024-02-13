package db

import (
	"fmt"
	"log"
	"main/pkg/pipelines"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func handleMigrateError(modelName string, err error) {
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
	// Migrate the schema
	if false {
		handleMigrateError("Pipelines", tx.AutoMigrate(&pipelines.Pipeline{}))
		handleMigrateError("Stages", tx.AutoMigrate(&pipelines.Stage{}))
		handleMigrateError("StageLabels", tx.AutoMigrate(&pipelines.StageLabel{}))
		handleMigrateError("StageLabel", tx.AutoMigrate(&pipelines.StageLabel{}))
		handleMigrateError("Leads", tx.AutoMigrate(&pipelines.Lead{}))
		handleMigrateError("CustomField", tx.AutoMigrate(&pipelines.CustomField{}))
		handleMigrateError("LeadQualifiers", tx.AutoMigrate(&pipelines.LeadQualifiers{}))
	}
	tx.Commit()
	fmt.Print("Database connected successfully...")

	fmt.Println("DB Connection Successful...")

}
