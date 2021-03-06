# `go-employees-service`

Implementing CRUD-like operations to manage employees.

## Functional requirements

- List employees
- Get an employee
- Create employee
- Update employee
- Delete employee

## Non-functional requirements

- [Docker](./Dockerfile)
- SQL Migrations with [Goose](https://github.com/pressly/goose)
- ORM with [SQL Boiler](https://github.com/volatiletech/sqlboiler)
- Task management with [Task](https://github.com/go-task/task)

## How to run it?

### From the source code

- Navigate to the project root
- Execute the following command: `go run main.go`
- Start firing requests from: http://localhost:4000/employees

### Using docker

- Build and start: `docker-compose up`

## Routes

### Get employees

`[GET] http://localhost:4000/employees`

### Get an employees

`[GET] http://localhost:4000/employees/{id}`

### Create employee

`[POST] http://localhost:4000/employees`

```
// Body
{
    "name": "John D. Smith",
    "role": "Software Engineer"
}
```
