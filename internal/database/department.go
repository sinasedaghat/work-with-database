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

func seedDepartment() error {
	query := `
		INSERT INTO department (name) 
		VALUES ($1) 
		ON CONFLICT (name) DO NOTHING
	`

	for _, departmentName := range departments {
		if _, err := DB.Exec(query, departmentName); err != nil {
			return err
		}
	}
	return nil
}
