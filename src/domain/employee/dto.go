package employee

import "time"

type CreateEmployeeDto struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	Country   string
}