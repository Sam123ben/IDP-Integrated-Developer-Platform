package main

import (
	"company_service/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/Sam123ben/IDP-Integrated-Developer-Platform/infra_env_dashboard/backend/common"
)

func main() {
	common.InitDB()         // Initialize the database connection
	defer common.DB.Close() // Close the database when done

	routes.SetupRoutes() // Set up routes

	fmt.Println("Company service is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
