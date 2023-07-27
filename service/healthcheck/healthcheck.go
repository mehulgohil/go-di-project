package healthecheck

import (
	"github.com/mehulgohil/go-di-project/interfaces"
	"github.com/mehulgohil/go-di-project/models"
)

type HealthCheckService struct {
	db interfaces.IDatabase
}

func (hc *HealthCheckService) CheckHealthCheck() (models.HealthCheckResponse, error) {
	return models.HealthCheckResponse{
		Status: "ok",
	}, nil
}

func NewHealthCheckService(db interfaces.IDatabase) *HealthCheckService {
	return &HealthCheckService{db: db}
}
