package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employeesList []Employee

func init() {
	employeesList = make([]Employee, 0)
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := json.Marshal(findEmployees())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Write(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "employees/")
	employeeId := urlPathSegments[len(urlPathSegments)-1]

	employee, err := findEmployee(employeeId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(employee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Write(response)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	err = json.Unmarshal([]byte(body), &employee)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	employee = saveEmployee(employee)
	response, err := json.Marshal(employee)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	employeesList = append(employeesList, employee)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

func employeesHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEmployees(w, r)
		return

	case http.MethodPost:
		createEmployee(w, r)
		return

	default:
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
	}
}

func employeeHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEmployee(w, r)
		return

	default:
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
	}
}

func main() {
	Port := 4000
	Host := "localhost"
	SrvAddress := Host + ":" + strconv.Itoa(Port)

	// Route handlers
	http.Handle("/employees", middleware(http.HandlerFunc(employeesHandlers)))
	http.Handle("/employees/", middleware(http.HandlerFunc(employeeHandlers)))

	fmt.Println("Server is up and running at", SrvAddress)
	err := http.ListenAndServe(SrvAddress, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
