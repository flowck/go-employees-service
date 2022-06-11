package employee

import "github.com/google/uuid"

var employees []Employee

func createEmployee(newEmployee CreateEmployeeDto) {
	employees = append(employees, Employee{
		Id:        uuid.NewString(),
		FirstName: newEmployee.FirstName,
		LastName:  newEmployee.LastName,
		BirthDate: newEmployee.BirthDate,
		Country:   newEmployee.Country,
	})
}
func readEmployees()           {}
func readEmployee()            {}
func updateEmployee()          {}
func deleteEmployee(id string) {}
