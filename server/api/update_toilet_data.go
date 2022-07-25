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

	id := json["id"]
	jsonPercentage := json["percentage"]
	percentage, _ := strconv.ParseFloat(fmt.Sprint(jsonPercentage), 32)
	location := json["location"]
	state := json["state"]
	toilet := model.Toilet{Id: fmt.Sprint(id), Percentage: float32(percentage), Location: fmt.Sprint(location), State: model.ToiletState(fmt.Sprint(state))}
	isExist := model.UpdateToiletData(toilet)

	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update toilet success!",
			"toilet":  model.GetToilet(fmt.Sprint(id)),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Toilet not exist!",
		})
	}

}
