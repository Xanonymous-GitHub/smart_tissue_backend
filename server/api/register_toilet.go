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
			"message": "Fail to register, please send the id of the toilet!",
		})
		return
	}

	restroomId, hasRestroomId := json["restroomId"]
	if !hasRestroomId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, please send the id of the restroom!",
		})
		return
	}

	isUndeployedToiletExist := model.IsUndeployedToiletExist(fmt.Sprint(toiletId))
	if !isUndeployedToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, toilet has been deployed or toilet not exist!",
		})
		return
	}

	isRestroomExist := model.IsRestroomExist(fmt.Sprint(restroomId))
	if !isRestroomExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, restroom not exist!",
		})
		return
	}

	model.RegisterToilet(fmt.Sprint(toiletId), fmt.Sprint(restroomId))

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register toilet success!",
		"toilet":  model.GetToilet(fmt.Sprint(toiletId)),
	})
}
