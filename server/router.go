package server

import (
	"backend/server/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Your server is working",
		})
	})

	router.GET("/restrooms", api.GetAllRestrooms)
	router.POST("/restroom", api.RegisterRestroom)

	_ = router.Run()
}
