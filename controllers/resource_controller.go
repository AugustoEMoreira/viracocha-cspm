package controllers

import (
	"github.com/AugustoEMoreira/viracocha-cspm/lib"
	"github.com/AugustoEMoreira/viracocha-cspm/services"
	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	logger  lib.Logger
	service services.ResourceService
}

func NewResourceController(s services.ResourceService, l lib.Logger) ResourceController {
	return ResourceController{
		service: s,
		logger:  l,
	}
}

func (r ResourceController) GetResource(c *gin.Context) {
	paramID := c.Param("id")

}
