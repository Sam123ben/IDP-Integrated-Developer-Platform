// backend/services/theme_service/cmd/main.go

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Initialize database connection
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	http.HandleFunc("/theme", themeHandler)
	log.Println("Theme service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func themeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTheme(w, r)
	case http.MethodPost:
		setTheme(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTheme(w http.ResponseWriter, r *http.Request) {
	var theme string
	err := db.QueryRow("SELECT theme FROM user_theme_preferences WHERE user_id = $1", 1).Scan(&theme) // assuming user_id=1 for demo
	if err != nil {
		http.Error(w, "Error fetching theme", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"theme": theme})
}

func setTheme(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Theme string `json:"theme"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	_, err := db.Exec("UPDATE user_theme_preferences SET theme = $1 WHERE user_id = $2", req.Theme, 1) // assuming user_id=1 for demo
	if err != nil {
		http.Error(w, "Error updating theme", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
