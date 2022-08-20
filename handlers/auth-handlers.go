package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/terawatthour/we-were-here-server/pkg/api"
)

type authHandlers struct {
	PostgresClient *sqlx.DB
}

type AuthHandlers interface {
	Mount(app *gin.Engine)
	googleEndpoint(c *gin.Context, ctx api.Context) error
}

func NewAuthHandlers(PostgresClient *sqlx.DB) AuthHandlers {
	return &authHandlers{PostgresClient: PostgresClient}
}

func (ctr authHandlers) Mount(app *gin.Engine) {
	auth := app.Group("api/v1/auth")
	auth.GET("/")
}

func (ctr authHandlers) googleEndpoint(c *gin.Context, ctx api.Context) error {
	return nil
}
