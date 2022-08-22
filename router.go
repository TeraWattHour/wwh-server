package main

import (
	"github.com/gorilla/mux"
	"github.com/terawatthour/we-were-here-server/handlers"
	"github.com/terawatthour/we-were-here-server/pkg/data"
)

func MountRoutes(r *mux.Router) {
	handlers.NewAuthHandlers(data.PostgresClient).Mount(r)
}
