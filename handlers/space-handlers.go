package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type spaceHandlers struct {
	PostgresClient *sqlx.DB
}

type SpaceHandlers interface {
	Mount(r *mux.Router)
}

func NewSpaceHandlers(PostgresClient *sqlx.DB) SpaceHandlers {
	return &spaceHandlers{PostgresClient: PostgresClient}
}

func (ctr spaceHandlers) Mount(r *mux.Router) {
	_ = r.NewRoute().PathPrefix("/api/v1/space").Subrouter()
}
