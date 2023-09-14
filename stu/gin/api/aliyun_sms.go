package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/stu/gin/global"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func ALiYunSms(c *gin.Context) {
	phone := c.Query("phone")
	rand.Seed(time.Now().Unix())
	codeInt := rand.Intn(1000000)
	data := global.Config.ALiYunSms.SendSMS(phone, strconv.Itoa(codeInt))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
}
