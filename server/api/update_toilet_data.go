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

	toiletId, hasToiletId := json["toiletId"]
	if !hasToiletId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the id of the target toilet!",
		})
		return
	}

	rawPercentage, hasPercentage := json["percentage"]
	if !hasPercentage {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, please send the new percentage of the toilet!",
		})
		return
	}
	percentage, error := strconv.ParseFloat(fmt.Sprint(rawPercentage), 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update, can not parse percentage to type float32!",
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

	toilet := model.Toilet{Id: fmt.Sprint(toiletId), Percentage: percentage, Location: fmt.Sprint(location), State: model.ToiletState(fmt.Sprint(state))}
	isToiletExist := model.IsToiletExists(toilet.Id)
	if !isToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Toilet not exist!",
		})
		return
	}

	model.UpdateToiletData(toilet)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update toilet success!",
		"toilet":  model.GetSingleToilet(fmt.Sprint(toiletId)),
	})
}
