package api

import (
	"air-pollution/api/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func bootstrap(
	logger log.Logger,

	stationsController *controllers.StationsController,
	listenerStationController *controllers.ListenerStationsController,
) *gin.Engine {
	r := gin.Default()
	r.GET("/stations", func(c *gin.Context) {
		ctx, err := stationsController.GetStations()
		httpContextToGin(c, logger, ctx, err)
	})

	r.POST("/listener/stations/:stationId", func(c *gin.Context) {
		stationID, _ := strconv.Atoi(c.Param("stationId"))
		ctx, err := listenerStationController.AddStation(stationID, controllers.AddStationRequestBody{})
		c.JSON(204, nil)
		httpContextToGin(c, logger, ctx, err)
	})

	r.DELETE("/listener/stations/:stationId", func(c *gin.Context) {
		stationID, _ := strconv.Atoi(c.Param("stationId"))
		ctx, err := listenerStationController.DeleteStation(stationID)
		httpContextToGin(c, logger, ctx, err)
	})

	r.GET("/listener/stations", func(c *gin.Context) {
		ctx, err := listenerStationController.GetStations()
		httpContextToGin(c, logger, ctx, err)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
	return r
}

func httpContextToGin(ginCtx *gin.Context, logger log.Logger, ctx *HttpContext, err error) {
	if err != nil {
		ginCtx.JSON(500, gin.H{
			"error": "Internal error",
		})
		logger.Println(err)
	}

	ginCtx.JSON(200, gin.H{
		"message": "pong",
	})
}
