package user

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/service"
	"codetube.cn/interface-api/interfaces"
	"codetube.cn/proto/service_user_login"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	*interfaces.ApiInterfaceTrait

	method       string
	request      interface{}
	response     *service_user_login.LoginResultResponse
	LoginHandler func(ctx context.Context, client service_user_login.UserLoginClient) func() (*service_user_login.LoginResultResponse, error) //注册服务调用函数
}

func ApiLogin(method string) *Login {
	return &Login{ApiInterfaceTrait: interfaces.NewApiInterfaceTrait(), method: method, request: nil, response: nil, LoginHandler: nil}
}

func (l *Login) Handler() {
	l.WithHandler(func(c *gin.Context) {
		userLoginClient, err := service.Client.UserLogin()
		if err != nil {
			l.Api.FailureWithStatusMessage(codes.ServiceConnectedFail, err.Error()).AbortWithStatusJSON(http.StatusInternalServerError)
			return
		}
		defer func() {
			if rc := recover(); rc != nil {
				l.Api.FailureWithStatusMessage(codes.ApiFailure, "Login failure").AbortWithStatusJSON(http.StatusInternalServerError)
			}
			return
		}()
		response, err := l.LoginHandler(c, userLoginClient)()
		if err != nil {
			l.Api.FailureWithStatus(codes.ServiceRequestFail).FailureWithMessage(err.Error()).Abort()
			//@todo 记录错误日志
			log.Println(err)
			return
		}
		l.response = response
	})
}

func (l *Login) Request() {
	l.WithRequest(func(c *gin.Context) {
		//登录方式，默认账号密码
		if l.method == "" {
			l.method = "username"
		}
		switch l.method {
		case "username":
			l.requestUsername(c)
		case "email":
			l.requestEmail(c)
		case "mobile":
			l.requestMobile(c)
		default:
			l.Api.FailureWithStatusMessage(codes.InvalidParam, "invalid Login type").Abort()
			return
		}
	})
}

func (l *Login) Response() {
	l.WithResponse(func(c *gin.Context) {
		l.Api.WithStatus(int(l.response.GetStatus())).
			WithMessage(l.response.GetMessage()).
			WithData(&map[string]string{"token": l.response.GetToken()}).
			Response()
	})
}

func (l *Login) requestUsername(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter username").Abort()
		return
	}
	if password == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter password").Abort()
		return
	}
	l.request = &service_user_login.LoginUsernameRequest{
		Username: username,
		Password: password,
	}
	l.LoginHandler = func(ctx context.Context, client service_user_login.UserLoginClient) func() (*service_user_login.LoginResultResponse, error) {
		return func() (*service_user_login.LoginResultResponse, error) {
			return client.Username(ctx, l.request.(*service_user_login.LoginUsernameRequest))
		}
	}
}

func (l *Login) requestEmail(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter email").Abort()
		return
	}
	if password == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter password").Abort()
		return
	}
	l.request = &service_user_login.LoginEmailRequest{
		Email:    email,
		Password: password,
	}
	l.LoginHandler = func(ctx context.Context, client service_user_login.UserLoginClient) func() (*service_user_login.LoginResultResponse, error) {
		return func() (*service_user_login.LoginResultResponse, error) {
			return client.Email(ctx, l.request.(*service_user_login.LoginEmailRequest))
		}
	}
}

func (l *Login) requestMobile(c *gin.Context) {
	mobile := c.PostForm("mobile")
	verifyCode := c.PostForm("verify_code")
	if mobile == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter mobile").Abort()
		return
	}
	if verifyCode == "" {
		l.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter verify_code").Abort()
		return
	}
	l.request = &service_user_login.LoginMobileRequest{
		Mobile:     mobile,
		VerifyCode: verifyCode,
	}
	l.LoginHandler = func(ctx context.Context, client service_user_login.UserLoginClient) func() (*service_user_login.LoginResultResponse, error) {
		return func() (*service_user_login.LoginResultResponse, error) {
			return client.Mobile(ctx, l.request.(*service_user_login.LoginMobileRequest))
		}
	}
}
