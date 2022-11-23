package routes

import (
	"github.com/AugustoEMoreira/viracocha-cspm/lib"
)

type IndexRoutes struct {
	logger  lib.Logger
	handler lib.RequestHandler
}

func (r IndexRoutes) Setup() {
	r.logger.Info("setting up route")
	api := r.handler.Gin.Group("/v1").Use()
	{
		api.GET("/index")
	}
}
