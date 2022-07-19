package api

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRestrooms(c *gin.Context) {
	restrooms := model.RestroomList

	c.JSON(http.StatusOK, gin.H{
		"restrooms": restrooms,
	})
}
