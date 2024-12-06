package main

import (
	"fmt"
	"server/infrastructure/app"
	"server/internal/core/handlers"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	app.InitEnvConfigs()
	app.InitLogger()
	app.PostgresConnect(&app.Config.Postgres)

	if app.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	engine := gin.New()

	// Initialize middlewares
	engine.Use(middlewares.JSONRecovery())
	engine.Use(middlewares.Logger(app.SLog))
	engine.Use(middlewares.RequestBody())
	engine.Use(middlewares.CORSMiddleware())

	// Initialize repositories

	// Initialize usecases

	// Initialize handlers
	handlers.NewAppHandler(engine)

	server := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.Port)
	app.SLog.Info("Running golang server")
	engine.Run(server)
}
