package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terawatthour/we-were-here-server/pkg"
	"github.com/terawatthour/we-were-here-server/pkg/data"
)

func init() {
	if err := pkg.LoadEnvironment(".env"); err != nil {
		log.Fatal("Error loading environment variables... ", err)
	}

	if err := data.EstablishPostgresConnection(); err != nil {
		log.Fatal("Error establishing postgres connection... ", err)
	}
}

func main() {
	app := gin.Default()

	MountRoutes(app)

	app.Run("localhost:8000")
}
