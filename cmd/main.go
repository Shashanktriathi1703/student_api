package main

import (
	"fmt"
	"log"
	"net/http"

	// "net/http"

	"github.com/Shashanktriathi1703/student-api/internal/config"
	"github.com/Shashanktriathi1703/student-api/internal/database"
	"github.com/Shashanktriathi1703/student-api/internal/handler"
	"github.com/Shashanktriathi1703/student-api/internal/repository"
	"github.com/Shashanktriathi1703/student-api/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Create database connection
	db := database.NewPostgresConnection(cfg)
	defer db.Close()


	// Create dependencies
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)


	//create router
	r := mux.NewRouter()
	
	// User routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	fmt.Println("Welcome to students Api")

	// CORS middleware
	handler := cors.Default().Handler(r)

	log.Println("Server starting on :8080")

	//main line yahi hai otherwise kuch chalega bhi nhi to, ye likhna to must hai
	log.Fatal(http.ListenAndServe(":8080", handler))
}
