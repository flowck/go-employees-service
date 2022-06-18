package employee

import "github.com/go-chi/chi/v5"

func EmployeesRoutes(r chi.Router) {
	r.Route("/employees", func(r chi.Router) {
		r.Get("/", getEmployees)
		r.Post("/", createEmployee)
	})

	r.Route("/employees/{employeeId}", func(r chi.Router) {
		r.Get("/", getEmployee)
	})
}
