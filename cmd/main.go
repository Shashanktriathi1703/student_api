package main

import (
	"fmt"
	"log"

	// "net/http"

	"github.com/Shashanktriathi1703/student-api/internal/config"
	"github.com/Shashanktriathi1703/student-api/internal/database"
	"github.com/Shashanktriathi1703/student-api/internal/repository"
	"github.com/Shashanktriathi1703/student-api/internal/service"
	"github.com/gorilla/mux"
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
	fmt.Println(userService)


	//create router
	r := mux.NewRouter()
	fmt.Println(r)

	fmt.Println("Welcome to students Api")

	log.Println("Server starting on :8080")
	// log.Fatal(http.ListenAndServe(":8080", handler))
}
