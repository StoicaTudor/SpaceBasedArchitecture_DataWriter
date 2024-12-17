package database_connector

import (
	"DataWriter/environment"
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	// Get connection details from environment variables or config file
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv(string(environment.DBUser)),
		os.Getenv(string(environment.DBPassword)),
		os.Getenv(string(environment.DBHost)),
		os.Getenv(string(environment.DBPort)),
		os.Getenv(string(environment.DBName)),
	)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return nil, err
	}

	return db, nil
}

//func InitDB() *sql.DB {
//	db, dbInitError := initDB()
//	if dbInitError != nil {
//		log.Printf("Error initializing database: %v", dbInitError)
//		return nil
//	}
//
//	return db
//}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing the database: %v", err)
		}
	}
}
