package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/simonalong/gole/config"
	http2 "github.com/simonalong/gole/http"
	"github.com/simonalong/gole/log"
	"github.com/simonalong/gole/util"
	"io/ioutil"
	"net/http"
	"strings"
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

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.Errorf("read request body failed,err = %s.", err)
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 状态码
		statusCode := c.Writer.Status()

		var body interface{}
		bodyStr := string(data)
		if "" != bodyStr {
			if strings.HasPrefix(bodyStr, "{") && strings.HasSuffix(bodyStr, "}") {
				var bodys []interface{}
				_ = util.StrToObject(bodyStr, &bodys)
				body = bodys
			} else if strings.HasPrefix(bodyStr, "[") && strings.HasSuffix(bodyStr, "]") {
				_ = util.StrToObject(bodyStr, &body)
			}
		}

		request := Request{
			Method:     c.Request.Method,
			Uri:        c.Request.RequestURI,
			Ip:         c.ClientIP(),
			Parameters: c.Params,
			Body:       body,
		}

		if config.GetValueBool("gole.show.head") {
			request.Headers = c.Request.Header
		}

		message := ErrorMessage{
			Request:    request,
			StatusCode: statusCode,
			Cost:       time.Now().Sub(startTime).String(),
		}

		if statusCode != 200 {
			for _, code := range exceptCode {
				if code == statusCode {
					return
				}
			}
			logger.WithField("result", util.ObjectToJson(message)).Error("请求异常")
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
					message.Response = response
					logger.WithField("result", util.ObjectToJson(message)).Error("请求异常")
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
	Body       interface{}
}

type ErrorMessage struct {
	Request    Request
	Response   http2.StandardResponse
	Cost       string
	StatusCode int
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
