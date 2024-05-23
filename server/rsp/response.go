package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResponse[T any] struct {
	ResponseBase
	Data T `json:"data"`
}

type DataArrayResponse[T any] struct {
	ResponseBase
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
	ResponseBase
	Data PagedData[T] `json:"data"`
}

func Success(ctx *gin.Context, object any) {
	ctx.JSON(http.StatusOK, object)
}

func SuccessOfStandard(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    0,
		"message": "success",
		"data":    v,
	})
}
