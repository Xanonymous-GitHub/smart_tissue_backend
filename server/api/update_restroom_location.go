package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestroomLocation(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please send the id of the restroom.",
		})
		return
	}

	location, hasLocation := json["location"]
	if !hasLocation {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please send the location of this restroom.",
		})
		return
	}

	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please register a restroom or request an existing restroom.",
		})
		return
	}

	model.UpdateRestroomLocation(fmt.Sprint(restroomId), fmt.Sprint(location))

	c.JSON(http.StatusOK, gin.H{
		"message":  "Update restroom location success!",
		"restroom": model.GetRestroom(fmt.Sprint(restroomId)),
	})
}
