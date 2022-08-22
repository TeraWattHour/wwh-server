package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terawatthour/we-were-here-server/pkg"
	"github.com/terawatthour/we-were-here-server/pkg/data"
)

func init() {
	if err := pkg.LoadEnvironment(".env"); err != nil {
		log.Fatal("Error loading environment variables... ", err)
	}

	if err := data.EstablishPostgresConnection(); err != nil {
		log.Fatal("Error establishing Postgres connection... ", err)
	}
}

func main() {
	r := mux.NewRouter()

	MountRoutes(r)

	http.ListenAndServe("localhost:8000", r)
	// app.Run("localhost:8000")
}
