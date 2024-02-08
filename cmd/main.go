package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler function to handle user registration
func registerUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to extract user data
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user data to the database or perform any necessary actions
	// For simplicity, let's just log the user data for now
	log.Printf("New user registered: %+v\n", user)

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	
	// Define paths to database and schema file
	dbPath := "db/initDB.go"
	schemaFile := "db/schema.sql"

	// Initialize the database
	db := db.InitializeDB(dbPath, schemaFile)

	// Close the database connection when the application exits
	defer db.Close()
	// Define HTTP routes
	http.HandleFunc("/register", registerUser)

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
