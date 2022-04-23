package user

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/service"
	"codetube.cn/interface-api/interfaces"
	"codetube.cn/proto/service_user_register"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Register struct {
	*interfaces.ApiInterfaceTrait

	request         interface{}
	response        *service_user_register.RegisterResultResponse
	registerHandler func(ctx context.Context, client service_user_register.UserRegisterClient) func() (*service_user_register.RegisterResultResponse, error) //注册服务调用函数
}

func ApiRegister() *Register {
	return &Register{ApiInterfaceTrait: interfaces.NewApiInterfaceTrait(), request: nil, response: nil, registerHandler: nil}
}

func (r *Register) Handler() {
	r.WithHandler(func(c *gin.Context) {
		userRegisterClient, err := service.Client.UserRegister()
		if err != nil {
			r.Api.FailureWithStatusMessage(codes.ServiceConnectedFail, err.Error()).AbortWithStatusJSON(http.StatusInternalServerError)
			return
		}
		defer func() {
			if rc := recover(); rc != nil {
				r.Api.FailureWithStatusMessage(codes.ApiFailure, "register failure").AbortWithStatusJSON(http.StatusInternalServerError)
			}
			return
		}()
		response, err := r.registerHandler(c, userRegisterClient)()
		if err != nil {
			r.Api.FailureWithStatus(codes.ServiceRequestFail).FailureWithMessage(err.Error()).Abort()
			//@todo 记录错误日志
			log.Println(err)
			return
		}
		r.response = response
	})
}

func (r *Register) Request() {
	r.WithRequest(func(c *gin.Context) {
		//注册方式，默认账号密码
		registerType := c.PostForm("type")
		if registerType == "" {
			registerType = "username"
		}
		switch registerType {
		case "username":
			r.requestUsername(c)
		case "email":
			r.requestEmail(c)
		case "mobile":
			r.requestMobile(c)
		default:
			r.Api.FailureWithStatusMessage(codes.InvalidParam, "invalid register type").Abort()
			return
		}
	})
}

func (r *Register) Response() {
	r.WithResponse(func(c *gin.Context) {
		r.Api.WithStatus(int(r.response.GetStatus())).
			WithMessage(r.response.GetMessage()).
			WithData(&map[string]string{"id": r.response.GetId()}).
			Response()
	})
}

func (r *Register) requestUsername(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter username").Abort()
		return
	}
	if password == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter password").Abort()
		return
	}
	r.request = &service_user_register.RegisterUsernameRequest{
		Username: username,
		Password: password,
	}
	r.registerHandler = func(ctx context.Context, client service_user_register.UserRegisterClient) func() (*service_user_register.RegisterResultResponse, error) {
		return func() (*service_user_register.RegisterResultResponse, error) {
			return client.Username(ctx, r.request.(*service_user_register.RegisterUsernameRequest))
		}
	}
}

func (r *Register) requestEmail(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter email").Abort()
		return
	}
	if password == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter password").Abort()
		return
	}
	r.request = &service_user_register.RegisterEmailRequest{
		Email:    email,
		Password: password,
	}
	r.registerHandler = func(ctx context.Context, client service_user_register.UserRegisterClient) func() (*service_user_register.RegisterResultResponse, error) {
		return func() (*service_user_register.RegisterResultResponse, error) {
			return client.Email(ctx, r.request.(*service_user_register.RegisterEmailRequest))
		}
	}
}

func (r *Register) requestMobile(c *gin.Context) {
	mobile := c.PostForm("mobile")
	verifyCode := c.PostForm("verify_code")
	if mobile == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter mobile").Abort()
		return
	}
	if verifyCode == "" {
		r.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter verify_code").Abort()
		return
	}
	r.request = &service_user_register.RegisterMobileRequest{
		Mobile:     mobile,
		VerifyCode: verifyCode,
	}
	r.registerHandler = func(ctx context.Context, client service_user_register.UserRegisterClient) func() (*service_user_register.RegisterResultResponse, error) {
		return func() (*service_user_register.RegisterResultResponse, error) {
			return client.Mobile(ctx, r.request.(*service_user_register.RegisterMobileRequest))
		}
	}
}
