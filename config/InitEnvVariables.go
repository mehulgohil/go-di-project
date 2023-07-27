package config

import "github.com/caarlos0/env/v8"

var EnvVariables EnvConfig

type EnvConfig struct {
	Environment                 string `env:"ENVIRONMENT" envDefault:"DEV"`
	AppPort                     string `env:"WEBSITES_PORT" envDefault:"8080"`
	PostgresSQLConnectionString string `env:"POSTGRES_SQL_CONNECTION_STRING,required"`
}

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables: " + err.Error())
	}
}
