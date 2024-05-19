package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SetUserRoutes(router *http.ServeMux) {
	router.HandleFunc("/user", getUserAll)  // Will handle all methods
	router.HandleFunc("GET /user/{id}", getUserById)
	router.HandleFunc("POST /user", createUser)
}

type user struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}


func getUserAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You have all users"))
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("You have user with id: " + id))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("User created successfully: %v", u)
	w.Write([]byte(response))
}
