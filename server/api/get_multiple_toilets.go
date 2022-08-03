package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMultipleToilets(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	id, hasId := json["id"]
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get, please send the id of the restroom.",
		})
		return
	}

	if !model.IsRestroomExists(fmt.Sprint(id)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get, please register a restroom or request an existing restroom.",
		})
		return
	}

	multipleToilets := model.GetToiletsFromRestroom(fmt.Sprint(id))

	c.JSON(http.StatusOK, gin.H{
		"toilets": multipleToilets,
	})
}
