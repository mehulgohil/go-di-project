package config

import "go.uber.org/zap"

var ZapLogger *zap.Logger

func InitializeLogger() {
	var err error

	if EnvVariables.Environment == "DEV" {
		ZapLogger, err = zap.NewDevelopment()
		if err != nil {
			panic("unable to initialize logger")
		}
		return
	}
	ZapLogger, err = zap.NewProduction()
	if err != nil {
		panic("unable to initialize logger")
	}
}
