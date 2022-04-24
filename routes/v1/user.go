package v1

import (
	"codetube.cn/interface-api/v1/user"
	"github.com/gin-gonic/gin"
)

func UserRoutesRegister(group *gin.RouterGroup) {
	// 用户注册
	userRegisterGroup := ApiRouter.Group(group, "/user/register")
	{
		ApiRouter.Post("", user.ApiRegister(), userRegisterGroup)
	}

	// 用户登录
	userLoginGroup := ApiRouter.Group(group, "/user/login")
	{
		ApiRouter.Post("", user.ApiLogin("username"), userLoginGroup)
		ApiRouter.Post("/account", user.ApiLogin("username"), userLoginGroup)
		ApiRouter.Post("/mobile", user.ApiLogin("mobile"), userLoginGroup)
		ApiRouter.Post("/email", user.ApiLogin("email"), userLoginGroup)
	}
}
