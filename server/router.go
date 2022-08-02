package server

import (
	"backend/server/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/restrooms", api.GetAllRestrooms)
	router.POST("/restroom", api.RegisterRestroom)
	router.POST("/toilet", api.RegisterToilet)
	router.PUT("/toilet", api.UpdateToiletData)
	router.DELETE("/toilet", api.RemoveToilet)

	_ = router.Run()
}
