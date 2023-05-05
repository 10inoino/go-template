package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (controller *HealthCheckController) HealthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
