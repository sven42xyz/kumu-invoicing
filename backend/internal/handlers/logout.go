package handlers

import (
	"net/http"

	"chapter42.de/m/internal/database"
	"chapter42.de/m/internal/response"
	"github.com/gorilla/sessions"
)

const (
	SessionClose int = -1
)

func LogoutHandler(db *database.MySQLDB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get session id from r
		session, err := store.Get(r, "session-cookie")
		if err != nil || session.IsNew {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		// Close session in session store
		// SessionClose = -1 -> instructs the store to close the given session
		session.Options.MaxAge = SessionClose
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Could not close the session", http.StatusInternalServerError)
			return
		}

		// Redirect user to / (root)
		RootHandler()

		// Return success response
		response.WriteResponse(w, "Logout successful", http.StatusOK)
		return
	}
}
