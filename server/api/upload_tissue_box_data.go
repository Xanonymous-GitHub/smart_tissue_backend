package api

import (
	"backend/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadTissueBoxData(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	id, hasId := json["id"]
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the id of this toilet.",
		})
	}

	state, hasState := json["state"]
	if !hasState {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the state of this toilet.",
		})
	}

	rawDistance, hasDistance := json["distance"]
	if !hasDistance {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the distance of this toilet.",
		})
	}
	distance, error := strconv.ParseFloat(fmt.Sprint(rawDistance), 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, cannot parse distance to float.",
		})
	}

	rawMaxDistance, hasMaxDistance := json["maxDistance"]
	if !hasMaxDistance {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, please send the max distance of this toilet.",
		})
	}
	maxDistance, error := strconv.ParseFloat(fmt.Sprint(rawMaxDistance), 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register, cannot parse max distance to float.",
		})
	}

	toilet := model.Toilet{Id: fmt.Sprint(id), State: model.ToiletState(fmt.Sprint(state)), Distance: distance, MaxDistance: maxDistance}

	model.UploadTissueBoxData(toilet)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Upload toilet data success!",
		"restroom": model.GetSingleToilet(toilet.Id),
	})
}
