package handlers

import (
	"log"
	"net/http"

	"chapter42.de/m/internal/database"
	"chapter42.de/m/internal/response"
	"chapter42.de/m/internal/xml"
	"github.com/gorilla/sessions"
)

func FileHandler(db *database.MySQLDB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-cookie")
		if err != nil || session.IsNew {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case http.MethodPost:
			peppol, err := uploadFile(w, r, session)
			if err != nil {
				log.Println(err.Error())
				response.WriteResponse(w, "Fehler", http.StatusInternalServerError)
				return
			}

			log.Println(peppol)

			response.WriteResponse(w, "Still brewing...", http.StatusTeapot)
			break
		}
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request, session *sessions.Session) (*xml.StandardBusinessDocument, error) {
	// Parse peppol bis xml
	peppol := &xml.StandardBusinessDocument{}
	err := peppol.PeppolDecoder(r.Body)
	if err != nil {
		return peppol, err
	}
	return peppol, err
}
