// internal/registration/register.go

package registration

import (
	"encoding/json"
	"net/http"

	"github.com/StefanCDev/Euro24/db"
	"golang.org/x/crypto/bcrypt"
)

// User represents a registered user
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUserHandler handles user registration requests
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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

	// Save user data to the database
	// Note: This assumes you have a function in your db package to save user data
	// Replace db.SaveUser with the actual function call in your project
	err = db.SaveUser(user)
	if err != nil {
		http.Error(w, "Failed to save user data", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
