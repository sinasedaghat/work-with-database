package database

func createDepartmentTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS department (
			id SERIAL PRIMARY KEY,
			name VARCHAR(20) UNIQUE NOT NULL
		)
	`
	_, err := DB.Exec(query)
	return err
}

func seedDepartment() {}
