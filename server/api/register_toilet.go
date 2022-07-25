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

	toiletId := json["toiletId"]
	location := json["location"]
	toilet := model.Toilet{Id: fmt.Sprint(toiletId), Percentage: 0.0, Location: fmt.Sprint(location), State: model.ToiletState("MAINTAINING")}
	restroomId := json["restroomId"]

	model.RegisterToilet(toilet, fmt.Sprint(restroomId))

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register toilet success!",
		"toilet":  model.GetToilet(fmt.Sprint(toiletId)),
	})
}
