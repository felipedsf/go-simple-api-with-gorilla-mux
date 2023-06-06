package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

var indexId uint32 = 2

var database = map[uint32]Customer{
	1: {ID: 1, Name: "Felipe", Role: "Admin", Email: "test@test.com", Phone: 553133333, Contacted: false},
	2: {ID: 2, Name: "Paula", Role: "Admin", Email: "ptest@test.com", Phone: 553122222, Contacted: false},
}

type Customer struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     uint64 `json:"phone"`
	Contacted bool   `json:"contacted"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{id}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", updateCustomer).Methods(http.MethodPut)
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods(http.MethodDelete)

	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", router)
}

func deleteCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		notFound(writer)
		return
	}

	if _, ok := database[uint32(id)]; ok {
		delete(database, uint32(id))
	} else {
		notFound(writer)
		return
	}

}

func updateCustomer(writer http.ResponseWriter, request *http.Request) {
	var uc Customer

	reqBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(reqBody, &uc)

	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		notFound(writer)
		return
	}

	if c, ok := database[uint32(id)]; ok {
		c.Name = uc.Name
		c.Email = uc.Email
		c.Role = uc.Role
		c.Phone = uc.Phone
		c.Contacted = uc.Contacted

		database[uint32(id)] = c
		json.NewEncoder(writer).Encode(c)
	} else {
		notFound(writer)
		return
	}
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		notFound(writer)
		return
	}

	if c, ok := database[uint32(id)]; ok {
		json.NewEncoder(writer).Encode(c)
	} else {
		notFound(writer)
		return
	}
}

func addCustomer(writer http.ResponseWriter, request *http.Request) {
	var newCustomer Customer

	reqBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(reqBody, &newCustomer)
	indexId = indexId + 1
	newCustomer.ID = indexId
	database[indexId] = newCustomer

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(newCustomer)
}

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	d := make([]Customer, 0, len(database))
	for _, customer := range database {
		d = append(d, customer)
	}

	json.NewEncoder(writer).Encode(d)
}

func notFound(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(map[string]string{})
}
