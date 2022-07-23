package api

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMultipleToilets(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	id := json["id"]
	toilets := model.GetMultipleToilets(fmt.Sprint(id))

	c.JSON(http.StatusOK, gin.H{
		"toilets": toilets,
	})
}
