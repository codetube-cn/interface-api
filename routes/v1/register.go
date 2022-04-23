package v1

import (
	"codetube.cn/interface-api/v1/user"
	"github.com/gin-gonic/gin"
)

func UserRegisterRegister(group *gin.RouterGroup) {
	userGroup := ApiRouter.Group(group, "/user/register")
	{
		ApiRouter.Post("", user.ApiRegister(), userGroup)
	}
}
