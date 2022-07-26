package api

import (
	"backend/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateToiletData(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	id, hasId := json["id"]
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the id of the target toilet!",
		})
		return
	}
	jsonPercentage, hasPercentage := json["percentage"]
	if !hasPercentage {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the new percentage of the toilet!",
		})
		return
	}
	location, hasLocation := json["location"]
	if !hasLocation {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the new location of the toilet!",
		})
		return
	}
	state, hasState := json["state"]
	if !hasState {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the new state of the toilet!",
		})
		return
	}
	percentage, _ := strconv.ParseFloat(fmt.Sprint(jsonPercentage), 32)
	toilet := model.Toilet{Id: fmt.Sprint(id), Percentage: float32(percentage), Location: fmt.Sprint(location), State: model.ToiletState(fmt.Sprint(state))}
	isExist := model.IsToiletExist(toilet.GetId())
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Toilet not exist!",
		})
		return
	}

	model.UpdateToiletData(toilet)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update toilet success!",
		"toilet":  model.GetToilet(fmt.Sprint(id)),
	})
}
