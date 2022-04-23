package components

import (
	"codetube.cn/interface-api/interfaces"
	"github.com/gin-gonic/gin"
)

var RouterEngine = gin.Default()

// ApiDefaultVersion 默认 API 版本号
// 在 ApiForcePrefixVersion == false 的情况下：
// 对应版本号的 API 可以使用 /version/path 和 /path 两种路径进行访问
// 在 URL 路径不变的情况下，可以用此直接升级默认版本的 API
var ApiDefaultVersion = "v1"

// ApiForcePrefixVersion 是否强行使用带版本号路径的路由
// 如果设置为 true，则所有 API 路径都强行带上版本号前缀
// 如果设置为 false，则默认版本将额外提供不带版本号前缀的路由路径
var ApiForcePrefixVersion = true

type router struct {
	version string
	groups  []*gin.RouterGroup
}

func NewRouter(version string) *router {
	//默认 path 带版本号的路由分组
	groups := []*gin.RouterGroup{
		RouterEngine.Group("/" + version),
	}
	//如果是默认版本，且未强制增加版本号前缀，则加入不带版本号的分组
	if version == ApiDefaultVersion && ApiForcePrefixVersion == false {
		groups = append(groups, RouterEngine.Group(""))
	}
	return &router{version: version, groups: groups}
}

func (r *router) Load(routes ...func(group *gin.RouterGroup)) {
	//加载各分组
	for _, g := range r.groups {
		for _, r := range routes {
			r(g)
		}
	}
}

func (r router) getApiHandlers(api interfaces.ApiInterface) []gin.HandlerFunc {
	api.Middleware()
	api.Request()
	api.Handler()
	api.Response()
	return api.GetHandlers()
}

func (r *router) Register(method, path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.Handle(method, path, r.getApiHandlers(api)...)
}

func (r *router) Get(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.GET(path, r.getApiHandlers(api)...)
}

func (r *router) Post(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.POST(path, r.getApiHandlers(api)...)
}

func (r *router) Patch(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.PATCH(path, r.getApiHandlers(api)...)
}

func (r *router) Delete(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.DELETE(path, r.getApiHandlers(api)...)
}

func (r *router) Option(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.OPTIONS(path, r.getApiHandlers(api)...)
}

func (r *router) Head(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.HEAD(path, r.getApiHandlers(api)...)
}

func (r *router) Any(path string, api interfaces.ApiInterface, group *gin.RouterGroup) {
	group.Any(path, r.getApiHandlers(api)...)
}

func (r *router) Group(baseGroup *gin.RouterGroup, path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return baseGroup.Group(path, handlers...)
}
