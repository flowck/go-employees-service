package main

import (
	"errors"
	"github.com/google/uuid"
)

var data map[string]Employee

func init() {
	data = make(map[string]Employee)
}

func findEmployees() []Employee {
	list := make([]Employee, 0)

	for _, v := range data {
		list = append(list, v)
	}

	return list
}

func findEmployee(id string) (*Employee, error) {
	employee := data[id]

	if employee.Id != "" {
		return &employee, nil
	}

	return nil, errors.New("Not found.")
}

func saveEmployee(employee Employee) Employee {
	employee.Id = uuid.NewString()
	data[employee.Id] = employee

	return employee
}
