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
	must(createDepartmentTable, "departments table")
	must(createProgramTable, "programs table")
	must(createStudentTable, "students table")

	// Seed tables
	seed(seedDepartment, "department")
	seed(seedProgram, "department")
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
