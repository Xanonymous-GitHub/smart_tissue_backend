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

	id := model.GenerateNextToiletId()
	location := json["location"]
	toilet := model.Toilet{Id: id, Percentage: 0.0, Location: fmt.Sprint(location), State: model.ToiletState("MAINTAINING")} ///
	model.RegisterToilet(toilet)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register toilet success!",
		"toilet":  model.GetToilet(id),
	})
}
