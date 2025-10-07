package database

import (
	"fmt"
)

var programs = map[string][]string{
	"Computer Science": {"Algorithms and Data Structures"},
	"Mathematics":      {"Applied Linear Algebra"},
	"Physics":          {"Solid State Physics"},
	"Chemistry":        {"Organic Chemistry"},
	"Biology":          {"Molecular Genetics"},
	"History":          {"World Civilizations"},
	"Economics":        {"Microeconomic Theory"},
}

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

func seedProgram() error {
	for dep, prs := range programs {
		var depId int

		if err := DB.QueryRow(`SELECT id FROM department WHERE name = $1`, dep).Scan(&depId); err != nil {
			return fmt.Errorf("could not find department '%s': %w", dep, err)
		}

		for _, pr := range prs {
			if _, err := DB.Exec(`
				INSERT INTO program (name, department_id)
				VALUES ($1, $2)
				ON CONFLICT (name) DO NOTHING
			`, pr, depId); err != nil {
				return err
			}
		}
	}
	return nil
}
