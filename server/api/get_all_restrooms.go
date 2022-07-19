package api

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRestrooms(c *gin.Context) {
	restrooms := []model.Restroom{
		{Id: "id1", Location: "location1", ToiletIdList: []model.Toilet{}},
		{Id: "id2", Location: "location2", ToiletIdList: []model.Toilet{}},
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "This is a message just for debugging, remove it when it is unused.",
		"restrooms": restrooms,
	})
}
