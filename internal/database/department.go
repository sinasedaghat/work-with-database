package database

var departments = []string{
	"Computer Science",
	"Mathematics",
	"Physics",
	"Chemistry",
	"Biology",
	"History",
	"Economics",
}

func createDepartmentsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS departments (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) UNIQUE NOT NULL
		)
	`
	_, err := DB.Exec(query)
	return err
}

func seedDepartments() error {
	query := `
		INSERT INTO departments (name) 
		VALUES ($1) 
	`
	// TODO: i can't use this line my query `ON CONFLICT (name) DO NOTHING`

	for _, departmentName := range departments {
		if _, err := DB.Exec(query, departmentName); err != nil {
			return err
		}
	}
	return nil
}
