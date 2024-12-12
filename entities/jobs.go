package entities

import "time"

// Job represents the job entity.
type Job struct {
    Id          uint
    Title       string
    Category    Category
    Vacancies   uint
    Description string
    Salary      float64
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
