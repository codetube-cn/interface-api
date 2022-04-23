package user

//
//import (
//	"codetube.cn/core/codes"
//	"codetube.cn/core/service"
//	"codetube.cn/interface-api/interfaces"
//	"codetube.cn/proto/user/request"
//	"codetube.cn/proto/user/response"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//type GetUser struct {
//	*interfaces.ApiInterfaceTrait
//
//	request  *request.GetUserByIdRequest
//	response *response.UserResponse
//}
//
//func ApiGetUser() *GetUser {
//	return &GetUser{ApiInterfaceTrait: interfaces.NewApiInterfaceTrait(), request: nil, response: nil}
//}
//
//func (u *GetUser) Handler() {
//	u.WithHandler(func(c *gin.Context) {
//		userClient, err := service.Client.User()
//		if err != nil {
//			u.Api.FailureWithStatusMessage(codes.ServiceConnectedFail, err.Error()).AbortWithStatusJSON(http.StatusInternalServerError)
//			return
//		}
//		response, err := userClient.GetUserById(c, u.request)
//		if err != nil {
//			u.Api.FailureWithStatus(codes.ServiceRequestFail).Abort()
//			//@todo 记录错误日志
//			return
//		}
//		u.response = response
//	})
//}
//
//func (u *GetUser) Request() {
//	u.WithRequest(func(c *gin.Context) {
//		userid := c.Query("user_id")
//		if userid == "" {
//			u.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter user_id").Abort()
//			return
//		}
//		userId, err := strconv.Atoi(userid)
//		if err != nil {
//			u.Api.FailureWithStatusMessage(codes.InvalidParam, "invalid parameter user_id").Abort()
//			return
//		}
//		u.request = &request.GetUserByIdRequest{Id: int64(userId)}
//	})
//}
//
//func (u *GetUser) Response() {
//	u.WithResponse(func(c *gin.Context) {
//		u.Api.SuccessWithData(u.response).Response()
//	})
//}
