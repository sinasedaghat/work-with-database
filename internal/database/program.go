package database

import (
	"fmt"
)

var programs = map[string][]string{
	"Economics":        {"Microeconomic Theory"},
	"History":          {"World Civilizations"},
	"Biology":          {"Molecular Genetics"},
	"Chemistry":        {"Organic Chemistry"},
	"Physics":          {"Solid State Physics"},
	"Mathematics":      {"Applied Linear Algebra"},
	"Computer Science": {"Algorithms and Data Structures"},
}

func createProgramsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS programs (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			department_id INT REFERENCES departments(id) ON DELETE RESTRICT,
			UNIQUE (name, department_id)
		)
	`
	_, err := DB.Exec(query)
	return err
}

func seedPrograms() error {
	for dep, prs := range programs {
		var depId int

		if err := DB.QueryRow(`SELECT id FROM departments WHERE name = $1`, dep).Scan(&depId); err != nil {
			return fmt.Errorf("could not find departments '%s': %w", dep, err)
		}

		for _, pr := range prs {
			if _, err := DB.Exec(`
				INSERT INTO programs (name, department_id)
				VALUES ($1, $2)
				ON CONFLICT (name, department_id) DO NOTHING
				`, pr, depId); err != nil {
				return err
			}
		}
	}
	return nil
}
