package handlers

import (
	"net/http"

	"chapter42.de/m/internal/database"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *database.MySQLDB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request parameters
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Retrieve user from database
		user, err := db.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Compare passwords
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
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
		session.Values["username"] = username
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}
}
