package user

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	users = append(users, User{ID: "1", Username: "admin", Password: "********"})

	router.HandleFunc("/", GetUsers).Methods("GET")
}