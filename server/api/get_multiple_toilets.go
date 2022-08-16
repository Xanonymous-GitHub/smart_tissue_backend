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

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get, please send the id of the restroom.",
		})
		return
	}

	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get, please register a restroom or request an existing restroom.",
		})
		return
	}

	multipleToilets := model.GetToiletsFromRestroom(fmt.Sprint(restroomId))

	c.JSON(http.StatusOK, gin.H{
		"toilets": multipleToilets,
	})
}
