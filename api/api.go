package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *gin.Context) SendMessage {
	user := c.Params.ByName("name")
	db["haha"] = "wawa"
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
