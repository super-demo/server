package handlers

import (
	"fmt"
	"net/http"
	"server/infrastructure/app"
	"server/version"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler(r *gin.Engine) *AppHandler {
	handler := &AppHandler{}

	r.GET("/", handler.Info)

	return handler
}

func (h *AppHandler) Info(c *gin.Context) {
	version := fmt.Sprintf("%d.%d.%d",
		version.MajorVersion,
		version.MinorVersion,
		version.RevisionVersion,
	)

	c.JSON(http.StatusOK, gin.H{
		"service":     "server",
		"status":      "ok",
		"version":     version,
		"environment": app.Config.Environment,
	})
}
