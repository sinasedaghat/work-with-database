package database

func createProgramTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS program (
			id SERIAL PRIMARY KEY,
			name VARCHAR(20) NOT NULL,
			department_id INT REFERENCES department(id) ON DELETE RESTRICT
		)
	`
	_, err := DB.Exec(query)
	return err
}

func seedProgram() {}
