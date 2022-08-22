package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/terawatthour/we-were-here-server/pkg/api"
)

type authHandlers struct {
	PostgresClient *sqlx.DB
}

type AuthHandlers interface {
	Mount(r *mux.Router)
	googleCallback(w api.Writer, r *http.Request) error
}

func NewAuthHandlers(PostgresClient *sqlx.DB) AuthHandlers {
	return &authHandlers{PostgresClient: PostgresClient}
}

func (ctr authHandlers) Mount(r *mux.Router) {
	auth := r.NewRoute().PathPrefix("/api/v1/auth").Subrouter()
	auth.Handle("/google/callback", api.W(ctr.googleCallback))
}

func (ctr authHandlers) googleCallback(w api.Writer, r *http.Request) error {
	return nil
}
