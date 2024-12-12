package jobsmodel

import (
	"fp-pbkk2/config"
	"fp-pbkk2/entities"
	"database/sql"
	"log"
)

// GetAll retrieves all jobs from the database
func GetAll() []entities.Job {
	query := `
	SELECT 
		jobs.id,
		jobs.title,
		jobs.category_id,
		categories.name AS category_name,
		jobs.vacancies,
		jobs.description,
		jobs.salary,
		jobs.created_at,
		jobs.updated_at
	FROM jobs
	JOIN categories ON jobs.category_id = categories.id
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var jobs []entities.Job

	for rows.Next() {
		var job entities.Job
		var salary sql.NullFloat64

		err := rows.Scan(
			&job.Id,
			&job.Title,
			&job.Category.Id,
			&job.Category.Name,
			&job.Vacancies,
			&job.Description,
			&salary,
			&job.CreatedAt,
			&job.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}

		if salary.Valid {
			job.Salary = salary.Float64
		} else {
			job.Salary = 0
		}

		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return jobs
}

// Create inserts a new job into the database
func Create(job entities.Job) bool {
	result, err := config.DB.Exec(`
		INSERT INTO jobs(
			title, category_id, vacancies, description, salary, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		job.Title,
		job.Category.Id,
		job.Vacancies,
		job.Description,
		job.Salary,
		job.CreatedAt,
		job.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return LastInsertId > 0
}

// Detail retrieves a job by ID from the database
func Detail(id int) entities.Job {
	row := config.DB.QueryRow(`
	SELECT 
		jobs.id,
		jobs.title,
		jobs.category_id,
		categories.name AS category_name,
		jobs.vacancies,
		jobs.description,
		jobs.salary,
		jobs.created_at,
		jobs.updated_at
	FROM jobs
	JOIN categories ON jobs.category_id = categories.id
	WHERE jobs.id = ?`, id)

	var job entities.Job
	var salary sql.NullFloat64

	err := row.Scan(
		&job.Id,
		&job.Title,
		&job.Category.Id,
		&job.Category.Name,
		&job.Vacancies,
		&job.Description,
		&salary,
		&job.CreatedAt,
		&job.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	if salary.Valid {
		job.Salary = salary.Float64
	} else {
		job.Salary = 0
	}

	return job
}

// Update updates a job in the database by ID
func Update(id int, job entities.Job) bool {
	result, err := config.DB.Exec(`
	UPDATE jobs SET
		title = ?,
		category_id = ?,
		vacancies = ?,
		description = ?,
		salary = ?,
		updated_at = ?
	WHERE id = ?`,
		job.Title,
		job.Category.Id,
		job.Vacancies,
		job.Description,
		job.Salary,
		job.UpdatedAt,
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected > 0
}


func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM jobs WHERE id = ?", id)
	return err
}
