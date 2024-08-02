package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/isuraem/hex/internal/adapters/app/api"
	"github.com/isuraem/hex/internal/adapters/core/user"
	"github.com/isuraem/hex/internal/adapters/framework/left/db"
	rightDB "github.com/isuraem/hex/internal/adapters/framework/right/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Update the connection string format
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Remove the "postgres" argument
	adapter, err := rightDB.NewAdapter(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer adapter.CloseDbConnection()

	userDB := db.NewUserDB(adapter)
	userService := user.NewUserService(userDB)
	userAPI := api.NewUserAPI(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userAPI.AddUser).Methods("POST")
	r.HandleFunc("/users/{id}", userAPI.RemoveUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", userAPI.ViewUser).Methods("GET")
	r.HandleFunc("/users", userAPI.ListUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
