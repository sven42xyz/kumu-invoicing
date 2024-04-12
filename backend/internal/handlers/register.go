package handlers

import (
	"net/http"

	"chapter42.de/m/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(db *database.MySQLDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request parameters
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		// Insert user into database
		err = db.CreateUser(username, string(hashedPassword))
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered successfully"))
	}
}
