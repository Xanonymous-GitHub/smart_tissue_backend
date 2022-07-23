package server

import (
	"backend/server/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/restrooms", api.GetAllRestrooms)
	router.POST("/restroom", api.RegisterRestroom)
	router.GET("/toilets", api.GetMultipleToilets)

	_ = router.Run()
}
