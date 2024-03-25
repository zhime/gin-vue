package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/server/model/common/response"
)

type UserRouter struct {
}

func (b *BaseApi) GetUserInfo(c *gin.Context) {
	//uuid := utils.GetUserUuid(c)
	//ReqUser, err := userService.GetUserInfo(uuid)
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//	return
	//}
	//response.OkWithDetailed(gin.H{"userInfo": "ReqUser"}, "获取成功", c)
	response.OkWithDetailed(gin.H{"userInfo": "ReqUserTest"}, "获取成功", c)
}
