package api

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRestroom(c *gin.Context) {
	restroom := model.Restroom{}
	c.BindJSON(&restroom)

	if restroom.Location == "" {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := model.GenerateNextRestroomId()
	restroom.Id = id
	model.RegisterRestroom(restroom)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Register restroom success!",
		"restroom": model.GetRestroom(id),
	})
}
