package database

import "log"

func PQInitialize() {
	// Create Enums and Types
	must(createStudentStatusEnums, "student status enum")

	// Create tables
	must(createDepartmentTable, "departments table")
	must(createProgramTable, "programs table")
	must(createStudentTable, "students table")

	// Seed tables
	seedDepartment()
	seedProgram()
}

func must(createFunc func() error, statementRole string) {
	if err := createFunc(); err != nil {
		log.Fatalf("🗄️ Failed to create %s: %v", statementRole, err)
	}
}
