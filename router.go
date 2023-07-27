package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mehulgohil/go-di-project/config"
	"github.com/mehulgohil/go-di-project/interfaces"
)

func InitRouter(db interfaces.IDatabase) *gin.Engine {
	if config.EnvVariables.Environment != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	healthCheckController := NewServiceInjector().InjectHealthCheckController(db)

	app.GET("/healthcheck", healthCheckController.CheckServerHealthCheck)

	return app
}
