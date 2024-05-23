package test

import (
	"fmt"
	"github.com/simonalong/gole/isc"
	"github.com/simonalong/gole/listener"
	"github.com/simonalong/gole/server/rsp"
	"github.com/simonalong/gole/server/test/pojo"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/simonalong/gole/logger"
	"github.com/simonalong/gole/server"
)

func TestApiVersion(t *testing.T) {
	fmt.Printf("step 1\n")
	server.RegisterRouteWithHeaders("/api/sample", server.HmGet, []string{"isc-api-version"}, []string{"1.0"}, func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("hello 1.0"))
	})
	fmt.Printf("step 2\n")
	server.RegisterRouteWithHeaders("/api/sample", server.HmGet, []string{"isc-api-version"}, []string{"2.0"}, func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("hello 2.0"))
	})
	fmt.Printf("step 3\n")
	server.RegisterRouteWithHeaders("/api/sample", server.HmGet, []string{"isc-api-version"}, []string{"3.0"}, func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("hello 3.0"))
	})
	server.StartServer()
}

func TestErrorPrint(t *testing.T) {
	server.RegisterRoute("/api/data", server.HmGet, func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("hello 3.0"))
	})
	server.StartServer()
}

func TestServerGet(t *testing.T) {
	server.Get("/info", func(c *gin.Context) {
		logger.Debug("debug的日志")
		logger.Info("info的日志")
		logger.Warn("warn的日志")
		logger.Error("error的日志")
		c.Data(200, "text/plain", []byte("hello"))
	})

	// 测试事件监听机制
	listener.AddListener(listener.EventOfServerRunFinish, func(event listener.BaseEvent) {
		logger.Info("应用启动完成")
	})

	server.StartServer()
}

func TestServer2(t *testing.T) {
	server.Get("/test/req1", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("hello"))
	})

	server.Get("/test/req2", func(c *gin.Context) {
		rsp.SuccessOfStandard(c, "value")
	})

	server.Get("/test/req3/:key", func(c *gin.Context) {
		rsp.SuccessOfStandard(c, c.Param("key"))
	})

	server.Post("/test/rsp1", func(c *gin.Context) {
		testReq := pojo.TestReq{}
		_ = isc.DataToObject(c.Request.Body, &testReq)
		rsp.SuccessOfStandard(c, testReq)
	})

	server.Get("/test/err", func(c *gin.Context) {
		rsp.FailedOfStandard(c, 500, "异常")
	})

	server.Run()
}

func init() {
	// 添加服务器启动完成事件监听
	listener.AddListener(listener.EventOfServerRunFinish, func(event listener.BaseEvent) {
		logger.Info("应用启动完成")
	})

	// 添加服务器启动完成事件监听
	listener.AddListener(listener.EventOfServerStop, func(event listener.BaseEvent) {
		logger.Info("应用退出")
	})
}

func TestServerOnProfileIsPprof(t *testing.T) {
	server.Use(ApiVersionInterceptor())
	server.Get("data", func(c *gin.Context) {
		rsp.SuccessOfStandard(c, "data")
	})

	server.Run()
}

func ApiVersionInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
