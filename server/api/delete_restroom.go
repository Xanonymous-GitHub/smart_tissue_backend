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

	id, hasId := json["id"]
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete, please send the id of the restroom.",
		})
		return
	}

	if !model.IsRestroomExists(fmt.Sprint(id)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete, please replace with an existing id.",
		})
		return
	}

	model.DeleteRestroom(fmt.Sprint(id))

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete restroom success!",
	})
}
