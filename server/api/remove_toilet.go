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
			"message": "Fail to remove, please send the id of the toilet!",
		})
		return
	}
	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to remove, please send the id of the restroom!",
		})
		return
	}
	isToiletExist := model.IsToiletExists(fmt.Sprint(toiletId))
	if !isToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to remove, toilet not exist or has not been registered!",
		})
		return
	}
	isRestroomExist := model.IsRestroomExists(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to remove, restroom not exist!",
		})
		return
	}
	isToiletIdInRestroom := model.IsToiletIdInRestroom(fmt.Sprint(toiletId), fmt.Sprint(restroomId))
	if !isToiletIdInRestroom {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to remove, toilet is not belong to this restroom!",
		})
		return
	}

	model.RemoveToilet(fmt.Sprint(toiletId), fmt.Sprint(restroomId))
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete toilet success!",
	})
}
