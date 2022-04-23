package course

//
//import (
//	"codetube.cn/core/codes"
//	"codetube.cn/core/service"
//	"codetube.cn/interface-api/interfaces"
//	"codetube.cn/proto/course/request"
//	"codetube.cn/proto/course/response"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//type GetCourses struct {
//	*interfaces.ApiInterfaceTrait
//
//	request  *request.GetCoursesRequest
//	response *response.CoursesResponse
//}
//
//func ApiGetCourses() *GetCourses {
//	return &GetCourses{ApiInterfaceTrait: interfaces.NewApiInterfaceTrait(), request: nil, response: nil}
//}
//
//func (this *GetCourses) Handler() {
//	this.WithHandler(func(ctx *gin.Context) {
//		courseClient, err := service.Client.Course()
//		if err != nil {
//			this.Api.FailureWithStatusMessage(codes.ServiceConnectedFail, err.Error()).AbortWithStatusJSON(http.StatusInternalServerError)
//			return
//		}
//		log.Println(courseClient)
//		response, err := courseClient.GetCourses(ctx, this.request)
//		if err != nil {
//			this.Api.FailureWithStatus(codes.ServiceRequestFail).Abort()
//			//@todo 记录错误日志
//			return
//		}
//		this.response = response
//	})
//}
//
//func (this *GetCourses) Request() {
//	this.WithRequest(func(ctx *gin.Context) {
//		var requestPage = int64(1)
//		var requestPageSize = int64(15)
//		queryPage := ctx.Query("page")
//		if queryPage != "" {
//			page, err := strconv.Atoi(queryPage)
//			if err != nil || page < 1 {
//				this.Api.FailureWithStatusMessage(codes.InvalidParam, "invalid parameter page").Abort()
//				return
//			} else {
//				requestPage = int64(page)
//			}
//		}
//		queryPageSize := ctx.Query("page_size")
//		if queryPageSize != "" {
//			pageSize, err := strconv.Atoi(queryPageSize)
//			if err != nil || pageSize < 1 {
//				this.Api.FailureWithStatusMessage(codes.InvalidParam, "invalid parameter page_size").Abort()
//				return
//			} else {
//				requestPageSize = int64(pageSize)
//			}
//		}
//
//		this.request = &request.GetCoursesRequest{Page: requestPage, PageSize: requestPageSize}
//	})
//}
//
//func (this *GetCourses) Response() {
//	this.WithResponse(func(ctx *gin.Context) {
//		this.Api.SuccessWithData(this.response).Response()
//	})
//}
