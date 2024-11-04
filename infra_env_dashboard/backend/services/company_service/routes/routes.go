package routes

import (
	"company_service/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/api/company", controllers.GetCompanyName)
}
