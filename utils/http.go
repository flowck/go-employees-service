package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type GenericResponse struct {
	Message string `json:"message"`
}

func NewGenericResponse(message string) []byte {
	response, err := json.Marshal(GenericResponse{Message: message})

	if err != nil {
		panic(err)
	}

	return response
}

func NewInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(NewGenericResponse("Internal Server Error"))
	log.Fatal(err)
}
