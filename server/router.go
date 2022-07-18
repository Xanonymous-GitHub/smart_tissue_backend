package server

import (
	"backend/server/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Your server is working",
		})
	})

	router.GET("/restrooms", api.GetAllRestrooms)

	_ = router.Run()
}
