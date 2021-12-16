package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	http2 "github.com/simonalong/tools/http"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
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
func ResponseHandler() gin.HandlerFunc {
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

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		if statusCode != 200 {
			logger.WithFields(logrus.Fields{"code": statusCode, "method": reqMethod, "uri": reqUri, "costTime": costTime, "ip": clientIP}).Error("请求异常")
		} else {
			var response http2.StandardResponse
			err := json.Unmarshal([]byte(blw.body.String()), &response)
			if err != nil {
				logger.Errorf("err: %v", err.Error())
			} else {
				if response.Code != 0 && response.Code != 200 && response.Code != "200" && response.Code != "success" {
					logger.WithFields(logrus.Fields{"code": statusCode, "method": reqMethod, "uri": reqUri, "costTime": costTime, "ip": clientIP, "errMsg": response.Message}).Error("请求异常")
				}
			}
		}
	}
}
