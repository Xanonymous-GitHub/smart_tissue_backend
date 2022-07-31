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

	id, hasId := json["id"]
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please send the id of the restroom.",
		})
		return
	}

	newLocation, hasLocation := json["location"]
	if !hasLocation {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please send the new location of this restroom.",
		})
		return
	}

	if !model.IsRestroomExists(fmt.Sprint(id)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, please register a restroom or request an existing restroom.",
		})
		return
	}

	model.UpdateRestroomLocation(fmt.Sprint(id), fmt.Sprint(newLocation))

	c.JSON(http.StatusOK, gin.H{
		"message":  "Update restroom location success!",
		"restroom": model.GetRestroom(fmt.Sprint(id)),
	})
}