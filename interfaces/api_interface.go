package interfaces

import (
	"codetube.cn/core/codes"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiInterface interface {
	GetHandlers() []gin.HandlerFunc // 组装函数

	Middleware() // 中间件函数，用于加入中间件
	Request()    // 请求处理函数
	Handler()    // 业务处理函数
	Response()   // 响应处理函数
}

type ApiInterfaceTrait struct {
	Handlers     []gin.HandlerFunc
	Middlewares  []gin.HandlerFunc
	RequestFunc  gin.HandlerFunc
	HandlerFunc  gin.HandlerFunc
	ResponseFunc gin.HandlerFunc
	Api          *ApiResponse
}

func NewApiInterfaceTrait() *ApiInterfaceTrait {
	return &ApiInterfaceTrait{}
}

// WithMiddleware 添加中间件函数
func (t *ApiInterfaceTrait) WithMiddleware(m ...gin.HandlerFunc) *ApiInterfaceTrait {
	t.Middlewares = append(t.Middlewares, m...)
	return t
}

// Middleware 设置中间件函数，默认什么也不做
func (t *ApiInterfaceTrait) Middleware() {
	//t.WithMiddleware(func(c *gin.Context) {})
}

// Request 请求函数，默认什么也不做
func (t *ApiInterfaceTrait) Request() {
	//t.WithRequest(func(c *gin.Context) {})
}

// Handler 默认业务处理函数，什么也不做
func (t *ApiInterfaceTrait) Handler() {
	//t.WithHandler(func(c *gin.Context) {})
}

// Response 响应函数，默认什么也不做
func (t *ApiInterfaceTrait) Response() {
	//t.WithResponse(func(c *gin.Context) {})
}

// WithRequest 添加 Request 请求处理函数
func (t *ApiInterfaceTrait) WithRequest(f gin.HandlerFunc) *ApiInterfaceTrait {
	t.RequestFunc = f
	return t
}

// WithHandler 添加业务处理函数
func (t *ApiInterfaceTrait) WithHandler(f gin.HandlerFunc) *ApiInterfaceTrait {
	t.HandlerFunc = f
	return t
}

// WithResponse 添加 Response 响应处理函数
func (t *ApiInterfaceTrait) WithResponse(f gin.HandlerFunc) *ApiInterfaceTrait {
	t.ResponseFunc = f
	return t
}

func (t *ApiInterfaceTrait) AppendHandler(handlerFunc ...gin.HandlerFunc) {
	t.Handlers = append(t.Handlers, handlerFunc...)
}

func (t *ApiInterfaceTrait) GetHandlers() []gin.HandlerFunc {
	// 加载公共组件
	// 响应组件
	t.AppendHandler(func(c *gin.Context) {
		t.Api = NewApiResponse(c)
	})

	// 加载业务组件
	// 先加载中间件
	if len(t.Middlewares) > 0 {
		t.AppendHandler(t.Middlewares...)
	}
	// 加载 request
	if t.RequestFunc != nil {
		t.AppendHandler(t.RequestFunc)
	}
	// 加载 handler
	if t.HandlerFunc != nil {
		t.AppendHandler(t.HandlerFunc)
	}
	// 加载 response
	if t.ResponseFunc != nil {
		t.AppendHandler(t.ResponseFunc)
	}
	return t.Handlers
}

// ApiResponse API 响应数据处理
type ApiResponse struct {
	status      int
	message     string
	data        interface{}
	context     *gin.Context
	defaultData interface{}
}

// NewApiResponse 实例化 ApiResponse
func NewApiResponse(c *gin.Context) *ApiResponse {
	return &ApiResponse{status: codes.Success, message: "success", defaultData: map[string]interface{}{}, context: c}
}

// WithStatus 设定响应状态识别码
func (r *ApiResponse) WithStatus(status int) *ApiResponse {
	r.status = status
	return r
}

// WithMessage 设定响应消息内容
func (r *ApiResponse) WithMessage(message string) *ApiResponse {
	r.message = message
	return r
}

// WithData 设定响应数据
func (r *ApiResponse) WithData(data interface{}) *ApiResponse {
	r.data = data
	return r
}

// Success 响应成功消息
func (r *ApiResponse) Success() *ApiResponse {
	return r.WithStatus(codes.Success).WithMessage("success").WithData(r.defaultData)
}

// SuccessWithMessage 响应带消息内容的成功信息
func (r *ApiResponse) SuccessWithMessage(message string) *ApiResponse {
	return r.WithStatus(codes.Success).WithMessage(message).WithData(r.defaultData)
}

// SuccessWithData 响应带数据的成功信息
func (r *ApiResponse) SuccessWithData(data interface{}) *ApiResponse {
	return r.WithStatus(codes.Success).WithMessage("success").WithData(data)
}

// SuccessWithMessageData 响应带消息内容和数据的成功消息
func (r *ApiResponse) SuccessWithMessageData(message string, data interface{}) *ApiResponse {
	return r.WithStatus(codes.Success).WithMessage(message).WithData(data)
}

// Failure 响应失败消息
func (r *ApiResponse) Failure() *ApiResponse {
	return r.WithStatus(codes.Failure).WithMessage("failure").WithData(r.defaultData)
}

// FailureWithStatus 响应带消息状态识别码的失败信息
func (r *ApiResponse) FailureWithStatus(status int) *ApiResponse {
	return r.WithStatus(status).WithMessage("failure").WithData(r.defaultData)
}

// FailureWithMessage 响应带消息内容的失败信息
func (r *ApiResponse) FailureWithMessage(message string) *ApiResponse {
	return r.WithStatus(codes.Failure).WithMessage(message).WithData(r.defaultData)
}

// FailureWithData 响应带数据的失败信息
func (r *ApiResponse) FailureWithData(data interface{}) *ApiResponse {
	return r.WithStatus(codes.Failure).WithMessage("failure").WithData(data)
}

// FailureWithMessageData 响应带消息内容和数据的失败消息
func (r *ApiResponse) FailureWithMessageData(message string, data interface{}) *ApiResponse {
	return r.WithStatus(codes.Failure).WithMessage(message).WithData(data)
}

// FailureWithStatusMessage 响应带状态识别码和消息内容的失败信息
func (r *ApiResponse) FailureWithStatusMessage(status int, message string) *ApiResponse {
	return r.WithStatus(status).WithMessage(message).WithData(r.defaultData)
}

// FailureWithStatusData 响应带状态识别码和消息内容的失败信息
func (r *ApiResponse) FailureWithStatusData(status int, data interface{}) *ApiResponse {
	return r.WithStatus(status).WithMessage("failure").WithData(data)
}

// Response 响应并输出
func (r *ApiResponse) Response() {
	r.ResponseWithStatus(http.StatusOK)
}

// ResponseWithStatus 指定 http status 进行响应并输出
func (r *ApiResponse) ResponseWithStatus(httpStatus int) {
	r.context.JSON(httpStatus, gin.H{
		"status":  r.status,
		"message": r.message,
		"data":    r.data,
	})
}

// Abort 响应并忽略后续函数
func (r *ApiResponse) Abort() {
	r.AbortWithStatusJSON(http.StatusOK)
}

// AbortWithoutData 无数据响应并忽略后续函数
func (r *ApiResponse) AbortWithoutData() {
	r.context.Abort()
}

// AbortWithStatus 指定 http status 响应并忽略后续函数
func (r *ApiResponse) AbortWithStatus(httpStatus int) {
	r.context.AbortWithStatus(httpStatus)
}

// AbortWithError 响应错误并忽略后续函数
func (r *ApiResponse) AbortWithError(httpStatus int, err error) {
	r.context.AbortWithError(httpStatus, err)
}

// AbortWithStatusJSON 指定 http status 响应数据并忽略后续函数
func (r *ApiResponse) AbortWithStatusJSON(httpStatus int) {
	r.context.AbortWithStatusJSON(httpStatus, gin.H{
		"status":  r.status,
		"message": r.message,
		"data":    r.data,
	})
}
