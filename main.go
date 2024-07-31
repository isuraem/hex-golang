// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/isuraem/hex/internal/adapters/app/api"
// 	"github.com/isuraem/hex/internal/adapters/core/arithmetic"
// 	"github.com/isuraem/hex/internal/adapters/framework/right/db"
// 	"github.com/isuraem/hex/internal/ports"
// 	"github.com/joho/godotenv"

// 	gRPC "github.com/isuraem/hex/internal/adapters/framework/left/grpc"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	if _, err := os.Stat(".env"); os.IsNotExist(err) {
// 		log.Fatalf(".env file does not exist")
// 	}

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	var dbaseAdapter ports.DBport
// 	var core ports.ArithmeticPort
// 	var appAdapter ports.APIPort
// 	var gRPCAdapter ports.GRPCPort

// 	dbaseDriver := os.Getenv("DB_DRIVER")
// 	dbaseHost := os.Getenv("DB_HOST")
// 	dbasePort := os.Getenv("DB_PORT")
// 	dbaseUser := os.Getenv("DB_USER")
// 	dbasePass := os.Getenv("DB_PASS")
// 	dbaseName := os.Getenv("DB_NAME")
// 	dbaseSSLMode := os.Getenv("DB_SSLMODE")

// 	// Print environment variables for debugging
// 	fmt.Printf("DB_DRIVER: %s\n", dbaseDriver)
// 	fmt.Printf("DB_HOST: %s\n", dbaseHost)
// 	fmt.Printf("DB_PORT: %s\n", dbasePort)
// 	fmt.Printf("DB_USER: %s\n", dbaseUser)
// 	fmt.Printf("DB_PASS: %s\n", dbasePass)
// 	fmt.Printf("DB_NAME: %s\n", dbaseName)
// 	fmt.Printf("DB_SSLMODE: %s\n", dbaseSSLMode)

// 	dsourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
// 		dbaseUser, dbasePass, dbaseHost, dbasePort, dbaseName, dbaseSSLMode)

// 	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
// 	if err != nil {
// 		log.Fatalf("failed to initiate dbase connection: %v", err)
// 	}

// 	log.Println("Successfully connected to the database.")

// 	defer dbaseAdapter.CloseDbConnection()

// 	core = arithmetic.NewAdapter()
// 	appAdapter = api.NewAdapter(dbaseAdapter, core)

//		gRPCAdapter = gRPC.NewAdapter(appAdapter)
//		gRPCAdapter.Run()
//	}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/isuraem/hex/internal/adapters/core/user"
	"github.com/isuraem/hex/internal/adapters/framework/left/api"
	"github.com/isuraem/hex/internal/adapters/framework/left/db"

	rightDB "github.com/isuraem/hex/internal/adapters/framework/right/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	adapter, err := rightDB.NewAdapter("postgres", connStr)
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
