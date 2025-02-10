package main

import (
	"fmt"
	"server/infrastructure/app"
	"server/internal/core/handlers"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
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
	authenticationRepository := repositories.NewAuthenticationRepository()
	userRepository := repositories.NewUserRepository(app.PostgresClient)

	// Initialize usecases
	authUsecase := usecases.NewAuthenticationUsecase(userRepository, authenticationRepository)
	userUsecase := usecases.NewUserUsecase(userRepository)

	// Initialize handlers
	handlers.NewAppHandler(engine)
	handlers.NewAuthenticationHandler(engine, authUsecase)
	handlers.NewUserHandler(engine, userUsecase, middlewares.Jwt())

	server := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.Port)
	app.SLog.Info("Running golang server")
	engine.Run(server)
}
