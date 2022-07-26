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

	location, hasLocation := json["location"]
	if !hasLocation {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, please send the location of the toilet!",
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
	
	toilet := model.Toilet{Id: fmt.Sprint(toiletId), Percentage: 0.0, Location: fmt.Sprint(location), State: model.ToiletState("MAINTAINING")}

	model.RegisterToilet(toilet, fmt.Sprint(restroomId))

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register toilet success!",
		"toilet":  model.GetToilet(fmt.Sprint(toiletId)),
	})
}
