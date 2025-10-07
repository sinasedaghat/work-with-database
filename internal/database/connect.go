package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PQConnection(environmentPath string) {
	// Create Data Source Name
	dsn := pqDataSourceName(environmentPath)

	// Open database connection
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("🗄️ Error from PQConnection() function in database package when call sql.Open(): %v", err)
	}

	// Verify connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("🗄️ Error from PQConnection() function in database package when call DB.Ping(): %v", err)
	}

	// Connection pool tuning
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func pqDataSourceName(envPath string) string {
	// Load .env file from the given path
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("🗄️ Error from PQDataSourceName() function in database package when load .env file from %s: %v", envPath, err)
	}

	// read .env file
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("🗄️ Error from PQDataSourceName() function in database package because one of .env fields not valid.")
	}

	// create data source name
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
