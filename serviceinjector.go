package main

import (
	hcController "github.com/mehulgohil/go-di-project/controllers/healthcheck"
	"github.com/mehulgohil/go-di-project/interfaces"
	hcService "github.com/mehulgohil/go-di-project/service/healthcheck"
	"sync"
)

type ServiceInjector struct{}

func (s *ServiceInjector) InjectHealthCheckController(db interfaces.IDatabase) *hcController.Controller {
	// injecting service layer in controller
	return hcController.NewHealthCheckController(hcService.NewHealthCheckService(db))
}

var (
	serviceInjectorObj *ServiceInjector
	injectOnce         sync.Once
)

func NewServiceInjector() *ServiceInjector {
	injectOnce.Do(func() {
		serviceInjectorObj = &ServiceInjector{}
	})
	return serviceInjectorObj
}
