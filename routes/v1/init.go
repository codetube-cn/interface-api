package v1

import (
	"codetube.cn/interface-api/components"
	"github.com/gin-gonic/gin"
)

var apiVersion = "v1"

var ApiRouter = components.NewRouter(apiVersion)

// LoadRoutes 需要注册的路由
var LoadRoutes = []func(group *gin.RouterGroup){
	UserRegister,
	CourseRegister,
	UserRegisterRegister,
}
