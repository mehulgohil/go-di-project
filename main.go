package main

import (
	_ "github.com/lib/pq"
	"github.com/mehulgohil/go-di-project/config"
	"github.com/mehulgohil/go-di-project/database"
	_ "github.com/mehulgohil/go-di-project/docs"
)

// @title			Dependency Injection Demo Project
// @version		1.0
// @description	This is a demo application to show dependency injection in Golang
// @host			localhost:8080
// @BasePath		/
func main() {
	//initialize env variables
	config.LoadEnvVariables()

	//initialize logger
	config.InitializeLogger()

	//initialize db connection
	dbClient, err := database.NewPGClient(config.EnvVariables.PostgresSQLConnectionString)
	if err != nil {
		config.ZapLogger.Panic("unable to connect database: " + err.Error())
	}

	//initialize api routes
	app := InitRouter(dbClient)

	//initialize swagger routes
	config.InitSwaggerRouter(app)

	err = app.Run(":" + config.EnvVariables.AppPort)
	if err != nil {
		config.ZapLogger.Panic("unable to start backend server: " + err.Error())
	}
}
