package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	Name string `json:"name""`
	Role string `json:"role"`
}

var employeesList []Employee

func getEmployeesHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		employeesResponse, err := json.Marshal(employeesList)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(employeesResponse)

	case http.MethodPost:
		var newEmployee Employee
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		err = json.Unmarshal([]byte(body), &newEmployee)

		if err != nil {
			log.Fatal(err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		employeesList = append(employeesList, newEmployee)

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte("Success"))

	default:
		res.WriteHeader(404)
		res.Write([]byte("Not found"))
	}
}

func main() {
	Port := 4000
	Host := "localhost"
	SrvAddress := Host + ":" + strconv.Itoa(Port)

	// Route handlers
	http.HandleFunc("/employees", getEmployeesHandler)

	fmt.Println("Server is up and running at", SrvAddress)
	err := http.ListenAndServe(SrvAddress, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
