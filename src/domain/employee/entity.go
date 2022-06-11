package employee

import "time"

type Employee struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName""`
	LastName  string    `json:"lastName""`
	BirthDate time.Time `json:"birthDate"`
	Country   string    `json:"country"`
}
