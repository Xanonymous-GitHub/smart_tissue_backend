package api

import (
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRestroom(c *gin.Context) {
	id := strconv.Itoa(model.NextRestroomId)
	restroom := model.Restroom{
		Id: id,
	}
	c.BindJSON(&restroom)

	if restroom.Location == "" {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	model.RestroomList[id] = restroom
	model.NextRestroomId += 1

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Register restroom success!",
		"restroom": model.RestroomList[id],
	})
}
