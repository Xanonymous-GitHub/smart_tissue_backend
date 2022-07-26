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
	toiletId := json["toiletId"]
	restroomId := json["restroomId"]
	isExist := model.IsToiletExist(fmt.Sprint(toiletId))

	model.RemoveToilet(fmt.Sprint(toiletId), fmt.Sprint(restroomId))

	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete toilet success!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Toilet not exist!",
		})
	}

}
