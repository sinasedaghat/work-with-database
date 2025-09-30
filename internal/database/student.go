package database

func createStudentStatusEnums() error {
	// query := `
	// 	CREATE TYPE IF NOT EXISTS student_status AS ENUM (
	//   	'active',
	//   	'suspended',
	//   	'graduated',
	//   	'withdrawn'
	// 	)
	// `

	query := `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'student_status') THEN
				CREATE TYPE student_status AS ENUM (
					'active',
					'suspended',
					'graduated',
					'withdrawn'
				);
			END IF;
		END
		$$;
	`
	_, err := DB.Exec(query)
	return err
}

func createStudentTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS students (
			id BIGSERIAL PRIMARY KEY,
			student_number VARCHAR(20) UNIQUE NOT NULL,
			first_name VARCHAR(50) NOT NULL,
			last_name  VARCHAR(50) NOT NULL,
			identifier_code VARCHAR(20) UNIQUE NOT NULL,
			foreigner BOOLEAN NOT NULL DEFAULT FALSE,
			birth_date DATE,
			gender CHAR(1) CHECK (gender IN ('M','F','O')),
			email VARCHAR(100) UNIQUE NOT NULL,
			phone VARCHAR(20),

			department_id INT REFERENCES department(id) ON DELETE RESTRICT,
			program_id    INT REFERENCES program(id)    ON DELETE RESTRICT,

			enrollment_year SMALLINT,
			current_semester SMALLINT,
			status student_status DEFAULT 'active',

			grade DECIMAL(4,2) CHECK (grade >= 0 AND grade <= 20.00),

			address TEXT,
			emergency_contact_name VARCHAR(100),
			emergency_contact_phone VARCHAR(20),
			CHECK (
        (emergency_contact_name IS NULL AND emergency_contact_phone IS NULL)
        OR
        (emergency_contact_name IS NOT NULL AND emergency_contact_phone IS NOT NULL)
    	),

			created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := DB.Exec(query)
	return err
}
