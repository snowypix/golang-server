package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// users is a slice of User structs
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// handleGetUsers returns a list of all users
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// handleGetUser returns a single user by ID
func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// handleCreateUser creates a new user
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.ID = len(users) + 1
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}

// handleUpdateUser updates an existing user
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for i, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			var updateUser User
			err := json.NewDecoder(r.Body).Decode(&updateUser)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			users[i] = updateUser
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// handleDeleteUser deletes a user
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for i, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/users", handleGetUsers)
	http.HandleFunc("/user/", handleGetUser)
	http.HandleFunc("/users/create", handleCreateUser)
	http.HandleFunc("/users/update", handleUpdateUser)
	http.HandleFunc("/users/delete", handleDeleteUser)
	http.ListenAndServe(":8080", nil)
}
