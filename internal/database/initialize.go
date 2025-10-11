package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func PQInitialize() {
	// Create Enums and Types
	must(createStudentStatusEnums, "student status enum")

	// Create tables
	must(createDepartmentsTable, "departments table")
	must(createProgramsTable, "programs table")
	must(createStudentsTable, "students table")

	// Seed tables
	seed(seedDepartments, "departments")
	seed(seedPrograms, "programs")
}

func must(createFunc func() error, statementRole string) {
	if err := createFunc(); err != nil {
		log.Fatalf("🗄️ Failed to create %s: %v", statementRole, err)
	}
}

func seed(seedFunc func() error, tableName string) {
	if err := seedFunc(); err != nil {
		log.Fatalf("🌱 Failed to seed %s: %v", tableName, err)
	}
}
