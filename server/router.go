package server

import (
	"backend/server/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/restrooms", api.GetAllRestrooms)
	router.POST("/restroom", api.RegisterRestroom)
	router.DELETE("/toilet", api.RemoveToilet)

	_ = router.Run()
}
