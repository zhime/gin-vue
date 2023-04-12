package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ALiYunSms(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": "sent sms success",
	})
}
