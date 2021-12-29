package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	http2 "github.com/simonalong/tools/http"
	"github.com/simonalong/tools/log"
	"github.com/simonalong/tools/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 日志记录到文件
func ResponseHandler(exceptCode ...int) gin.HandlerFunc {
	//实例化
	logger := log.GetLogger("isc-config-service")

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		costTime := endTime.Sub(startTime)

		// 状态码
		statusCode := c.Writer.Status()

		bodyMap := map[string]interface{}{}
		_ = util.DataToObject(c.Request.Body, &bodyMap)

		request := Request{
			Method:     c.Request.Method,
			Uri:        c.Request.RequestURI,
			Ip:         c.ClientIP(),
			Headers:    c.Request.Header,
			Parameters: c.Params,
			Body:       bodyMap,
		}

		if statusCode != 200 {
			for _, code := range exceptCode {
				if code == statusCode {
					return
				}
			}
			logger.WithFields(logrus.Fields{"request": util.ObjectToJson(request), "costTime": costTime}).Error("请求异常")
		} else {
			var response http2.StandardResponse
			err := json.Unmarshal([]byte(blw.body.String()), &response)
			if err != nil {
				return
			} else {
				if response.Code == nil {
					return
				}
				if response.Code != 0 && response.Code != "0" && response.Code != 200 && response.Code != "200" && response.Code != "success" {
					logger.WithFields(logrus.Fields{"request": util.ObjectToJson(request), "response": util.ObjectToJson(response), "costTime": costTime}).Error("请求异常")
				}
			}
		}
	}
}

type Request struct {
	Method     string
	Uri        string
	Ip         string
	Headers    http.Header
	Parameters gin.Params
	Body       map[string]interface{}
}

func Success(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusOK, object)
}

func SuccessOfStandard(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    v,
	})
}

func FailedOfStandard(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}

func FailedWithDataOfStandard(ctx *gin.Context, code string, message string, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    v,
	})
}
