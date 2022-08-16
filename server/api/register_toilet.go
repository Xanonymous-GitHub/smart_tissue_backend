package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterToilet(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	toiletId, hasToiletId := json["toiletId"]
	if !hasToiletId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the id of the toilet",
		})
		return
	}

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the id of the restroom",
		})
		return
	}

	location, hasLocation := json["location"]
	if !hasLocation {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the location of the toilet",
		})
		return
	}

	isUndeployedToiletExist := model.IsUndeployedToiletExist(fmt.Sprint(toiletId))
	if !isUndeployedToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please request an undeployed toilet.",
		})
		return
	}

	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please register a restroom or request an existing restroom.",
		})
		return
	}

	model.RegisterToilet(fmt.Sprint(toiletId), fmt.Sprint(restroomId), fmt.Sprint(location))

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register toilet success!",
		"toilet":  model.GetSingleToilet(fmt.Sprint(toiletId)),
	})
}
