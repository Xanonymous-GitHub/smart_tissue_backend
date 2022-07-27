package api

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUndeployedToiletIds(c *gin.Context) {
	undeployedToiletIdList := model.GetUndeployedToiletIdList()

	c.JSON(http.StatusOK, gin.H{
		"undeployedToiletIds": undeployedToiletIdList,
	})
}
