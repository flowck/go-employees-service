package main

import (
	"employees-service/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func AppMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := models.Users(qm.Limit(5)).All(ctx, DB)

	employees, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Write(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := chi.URLParam(r, "employeeId")

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

type Config struct {
	DbDriver string `envconfig:"DB_DRIVER" required:"true"`
	DbUrl    string `envconfig:"DB_URL" required:"true"`
	PORT     int    `envconfig:"PORT" required:"true"`
	HOST     string `envconfig:"HOST" required:"true"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := Config{}
	err := envconfig.Process("", &config)
	SrvAddress := fmt.Sprintf("%s:%s", config.HOST, strconv.Itoa(config.PORT))
	r := chi.NewRouter()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	InitDatabase(&config)

	if err := goose.Up(DB, "sql"); err != nil {
		panic(err)
	}

	r.Use(middleware.Logger)
	r.Use(AppMiddleware)
	r.Route("/employees", func(r chi.Router) {
		r.Get("/", getEmployees)
		r.Post("/", createEmployee)
	})

	r.Route("/employees/{employeeId}", func(r chi.Router) {
		r.Get("/", getEmployee)
	})

	fmt.Println("Server is up and running at", SrvAddress)
	err = http.ListenAndServe(SrvAddress, r)

	if err != nil {
		log.Fatal(err)
		return
	}
}
