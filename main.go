package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type api struct {
	addr string
}

//	func (s *api) handleRequest(w http.ResponseWriter, r *http.Request) {
//		// fmt.Fprintf(w, "Hello, World!")
//		// w.write([]byte("Hello, World!"))
//		switch r.Method {
//		case http.MethodGet:
//			switch r.URL.Path {
//			case "/":
//				fmt.Fprintf(w, "Home Page")
//			case "/about":
//				fmt.Fprintf(w, "About Page")
//			default:
//				http.Error(w, "404 not found", http.StatusNotFound)
//			}
//		case http.MethodPost:
//			fmt.Fprintf(w, "POST Request")
//		default:
//			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		}
//	}

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
	users = append(users, u)

	// set header
	w.Header().Set("Content-Type", "application/json")

	// write the response
	w.WriteHeader(http.StatusCreated)
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
	mux.HandleFunc("GET /", api.getUserHandler)
	mux.HandleFunc("POST /", api.createUserHandler)

	fmt.Println("Server is running on port http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
