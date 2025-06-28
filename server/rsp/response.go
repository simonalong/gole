package rsp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/simonalong/gole/errorx"
	"net/http"
	"strings"
)

type ResponseBase struct {
	Code   string `json:"code"`
	Data   any    `json:"data,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type ResponseGole struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResponse[T any] struct {
	ResponseGole
	Data T `json:"data"`
}

type DataArrayResponse[T any] struct {
	ResponseGole
	Data []T `json:"data"`
}

type PagedData[T any] struct {
	Total         int64 `json:"total"`
	Size          int64 `json:"size"`
	Current       int64 `json:"current"`
	Pages         int64 `json:"pages"`
	IsSearchCount bool  `json:"isSearchCount"`
	Records       []T   `json:"records"`
}

type PagedResponse[T any] struct {
	ResponseGole
	Data PagedData[T] `json:"data"`
}

func Success(ctx *gin.Context, object any) {
	ctx.JSON(http.StatusOK, object)
}

func Fail(ctx *gin.Context, code int, object any) {
	ctx.JSON(code, object)
}

func SuccessOfStandard(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    0,
		"message": "success",
		"data":    v,
	})
}

func FailOfStandard(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}

func FailWithDataOfStandard(ctx *gin.Context, code string, message string, v any) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    code,
		"message": message,
		"data":    v,
	})
}

// Done 用于返回服务端响应处理结果。
// data 作为响应数据的封装对象，返回给前端。
// message 作为业务逻辑层处理错误时返回的错误信息对象。
// example：
//
//	rsp, err := rpc_clients.UserService.GetUser(context.Background(), req)
//	if err != nil {
//		Done(ctx, nil, err)
//		return
//	}
//	Done(ctx, rsp.User)
//
// if you want to Done a page list result:
//
//	data := response.ListWrap{Total: total, BsList: rsp.Users}
//	Done(ctx, data)
func Done(ctx *gin.Context, data any, errs ...error) {
	var xe *errorx.BaseError
	if len(errs) > 0 && errs[0] != nil {
		if !errors.As(errs[0], &xe) {
			xe = errorx.SC_SERVER_ERROR.WithError(errs[0])
		}
	}

	if xe == nil {
		xe = errorx.SC_OK
	}

	var body = ResponseBase{
		Code:   xe.Code,
		Msg:    xe.Msg,
		Detail: xe.Detail,
	}

	switch xe.Code {
	case "SC_OK", "SC_FOUND", "SC_MOVED_PERMANENTLY":
		body.Data = data
	}
	jsonResponse(ctx, &body, errorx.GetHttpStatus(xe.Code))
	return
}

func jsonResponse(ctx *gin.Context, body *ResponseBase, status int) {
	if strings.Contains(ctx.GetHeader("User-Agent"), "curl") {
		ctx.IndentedJSON(status, body)
		ctx.Abort()
	} else {
		ctx.AbortWithStatusJSON(status, body)
	}
}
