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

	rawMaxDistance, hasMaxDistance := json["maxDistance"]
	if !hasMaxDistance {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, please send the max distance of this toilet!",
		})
	}
	maxDistance, error := strconv.ParseFloat(fmt.Sprint(rawMaxDistance), 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to register, cannot parse max distance to float!",
		})
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

	isToiletExist := model.IsToiletExists(fmt.Sprint(toiletId))
	if !isToiletExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Toilet not exist!",
		})
		return
	}

	toilet := model.Toilet{Id: fmt.Sprint(toiletId), MaxDistance: maxDistance, Location: fmt.Sprint(location), State: model.ToiletState(fmt.Sprint(state))}

	model.UpdateToiletData(toilet)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update toilet success!",
		"toilet":  model.GetSingleToilet(fmt.Sprint(toiletId)),
	})
}
