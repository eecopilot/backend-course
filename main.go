package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// decode the request body
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// create a new user
	u := User{
		ID:   len(users) + 1,
		Name: payload.Name,
	}

	// insert the user
	if err := insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set header
	w.Header().Set("Content-Type", "application/json")

	// write the response
	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	// storage validation
	for _, user := range users {
		if user.Name == u.Name {
			return errors.New("user already exists")
		}
	}
	users = append(users, u)
	return nil
}

func main() {
	api := &api{addr: ":8080"}
	// initialize the serveMux
	mux := http.NewServeMux()

	// initialize the server
	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	// register the handler
	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	fmt.Println("Server is running on port http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		// log.Fatal(err)
		panic(err)
	}
}
