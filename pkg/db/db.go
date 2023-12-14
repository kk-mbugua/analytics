package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

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
	fmt.Print("Database connected successfully...")

	// TODO: AutoMigrate database models

	fmt.Println("DB Connection Successful...")

}
