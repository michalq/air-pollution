package api

import (
	"air-pollution/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeGin() *gin.Engine {
	wire.Build(
		bootstrap,

		controllers.NewStationsController,
		controllers.NewListenerStationsController,
	)

	return &gin.Engine{}
}
