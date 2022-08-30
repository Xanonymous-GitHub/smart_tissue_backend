package server

import (
	"backend/server/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.POST("/restroom", api.RegisterRestroom)
	router.GET("/restrooms", api.GetAllRestrooms)
	router.PUT("/restroom", api.UpdateRestroomLocation)
	router.DELETE("/restroom", api.DeleteRestroom)
	router.POST("/toilet", api.RegisterToilet)
	router.POST("/toilets", api.GetMultipleToilets)
	router.GET("/undeployedToiletIds", api.GetUndeployedToiletIds)
	router.PUT("/toilet", api.UpdateToiletData)
	router.DELETE("/toilet", api.RemoveToilet)
	router.PUT("/toiletData", api.UploadTissueBoxData)

	_ = router.Run()
}
