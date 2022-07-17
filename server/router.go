package server

import (
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

	_ = router.Run()
}
