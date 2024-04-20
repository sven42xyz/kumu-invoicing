package handlers

import (
	"net/http"

	"chapter42.de/m/internal/data"
	"chapter42.de/m/internal/database"
	"chapter42.de/m/internal/response"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *database.MySQLDB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request parameters
		data := &data.User{}
		data.JsonDecode(r.Body)

		if (len(data.Username) < 1 && len(data.Password) < 1) {
			http.Error(w, "Invalid arguments", http.StatusUnauthorized)
			return
		}

		// Retrieve user from database
		user, err := db.GetUserByUsername(data.Username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Compare passwords
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Authentication successful
		// Create session cookie
		session, err := store.New(r, "session-cookie")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["username"] = data.Username
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response
		response.WriteResponse(w, "Login successful", http.StatusOK)
		return
	}
}
