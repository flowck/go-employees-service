package employee

import (
	"employees-service/infra"
	"employees-service/models"
	"employees-service/utils"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
	"log"
	"net/http"
)

func getEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := models.Employees(qm.Limit(100)).All(ctx, infra.DB)

	employees, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Write(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	employeeId := chi.URLParam(r, "employeeId")

	employee, err := models.Employees(qm.Where("id = ?", employeeId)).One(ctx, infra.DB)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(utils.NewGenericResponse(fmt.Sprintf("User %s was not found.", employeeId)))
		return
	}

	response, err := json.Marshal(employee)

	if err != nil {
		utils.NewInternalServerError(w, err)
		return
	}

	w.Write(response)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload Employee
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(body, &payload)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newEmployee := models.Employee{FirstName: payload.FirstName, LastName: payload.LastName}
	err = newEmployee.Insert(ctx, infra.DB, boil.Infer())

	if err != nil {
		utils.NewInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(utils.NewGenericResponse("User created successfully"))
}
