package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terawatthour/we-were-here-server/handlers"
	"github.com/terawatthour/we-were-here-server/pkg/data"
)

func MountRoutes(app *gin.Engine) {
	handlers.NewAuthHandlers(data.PostgresClient).Mount(app)
}
