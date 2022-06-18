package main

import (
	"employees-service/infra"
	"employees-service/modules/employee"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"net/http"
	"strconv"
)

func AppMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	SrvAddress := fmt.Sprintf("%s:%s", infra.Configs.HOST, strconv.Itoa(infra.Configs.PORT))
	r := chi.NewRouter()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	infra.InitDatabase(&infra.Configs)

	if err := goose.Up(infra.DB, "sql"); err != nil {
		panic(err)
	}

	r.Use(middleware.Logger)
	r.Use(AppMiddleware)
	employee.EmployeesRoutes(r)

	fmt.Println("Server is up and running at", SrvAddress)
	err := http.ListenAndServe(SrvAddress, r)

	if err != nil {
		log.Fatal(err)
		return
	}
}
