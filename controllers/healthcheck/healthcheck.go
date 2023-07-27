package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/mehulgohil/go-di-project/interfaces"
	"github.com/mehulgohil/go-di-project/service/healthcheck"
	"net/http"
)

type Controller struct {
	interfaces.IHealthCheckService
}

// CheckServerHealthCheck checks the health status of the server.
// @Summary Check server health status
// @Description Checks the health status of the server
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} models.HealthCheckResponse
// @Failure 500 {object} map[string]string
// @Router /healthcheck [get]
func (controller *Controller) CheckServerHealthCheck(ctx *gin.Context) {
	check, err := controller.CheckHealthCheck()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, check)
}

func NewHealthCheckController(HealthCheckService *healthecheck.HealthCheckService) *Controller {
	return &Controller{
		HealthCheckService,
	}
}
