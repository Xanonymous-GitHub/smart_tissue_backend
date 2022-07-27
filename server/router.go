package server

import (
	"backend/server/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/restrooms", api.GetAllRestrooms)
	router.POST("/restroom", api.RegisterRestroom)
	router.PUT("/restroom", api.UpdateRestroomLocation)
	router.DELETE("/restroom", api.DeleteRestroom)
	router.GET("/toilets", api.GetMultipleToilets)
  router.GET("/undeployedToiletIds", api.GetUndeployedToiletIds)

	_ = router.Run()
}
