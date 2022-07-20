package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRestroom(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	id := model.GenerateNextRestroomId()
	location := json["location"]
	restroom := model.Restroom{Id: id, Location: fmt.Sprint(location), ToiletIdList: []string{}}
	model.RegisterRestroom(restroom)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Register restroom success!",
		"restroom": model.GetRestroom(id),
	})
}