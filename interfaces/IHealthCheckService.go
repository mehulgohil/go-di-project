package interfaces

import "github.com/mehulgohil/go-di-project/models"

type IHealthCheckService interface {
	CheckHealthCheck() (models.HealthCheckResponse, error)
}
