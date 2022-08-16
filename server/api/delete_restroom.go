package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestroom(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete, please send the id of the restroom.",
		})
		return
	}

	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete, please replace with an existing id.",
		})
		return
	}

	model.DeleteRestroom(fmt.Sprint(restroomId))

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete restroom success!",
	})
}
