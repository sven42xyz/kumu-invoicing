package handlers

import (
	"net/http"

	"chapter42.de/m/internal/database"
	"github.com/gorilla/sessions"
)

func UserHandler(db *database.MySQLDB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-cookie")
		if err != nil || session.IsNew {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case http.MethodPut:
			change(w, session)
			break
		case http.MethodDelete:
			delete(w, session)
			break
		}
	}
}

func change(w http.ResponseWriter, sessions *sessions.Session) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("UserChange not implemented"))
	return
}

func delete(w http.ResponseWriter, session *sessions.Session) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("UserDelete not implemented"))
	return
}
