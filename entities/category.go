package entities

import "time"

//mendefenisikan entities category
type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
