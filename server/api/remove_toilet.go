package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveToilet(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	toiletId, hasToiletId := json["toiletId"]
	if !hasToiletId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove, please send the id of the toilet.",
		})
		return
	}

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove, please send the id of the restroom.",
		})
		return
	}

	isToiletExist := model.IsToiletExists(fmt.Sprint(toiletId))
	if !isToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove, please request an existing toilet.",
		})
		return
	}

	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove, please request an existing restroom.",
		})
		return
	}

	isToiletIdInRestroom := model.IsToiletIdInRestroom(fmt.Sprint(toiletId), fmt.Sprint(restroomId))
	if !isToiletIdInRestroom {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove, please request a toilet which is belong to the restroom.",
		})
		return
	}

	model.RemoveToilet(fmt.Sprint(toiletId), fmt.Sprint(restroomId))
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete toilet success!",
	})
}
