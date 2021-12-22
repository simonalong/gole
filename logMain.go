package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simonalong/tools/config"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

var serviceLogger *logrus.Logger
var testLogger *logrus.Logger

func init() {
	log.LogConfig("/Users/zhouzhenyong/tem/tools/logs/tools", "/api/tools/", true)
}

func main() {
	//定义命令行参数方式1
	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	//
	////解析命令行参数
	//flag.Parse()
	//fmt.Println(name, age, married, delay)

	//config.LoadConfigWithPath("./test/resources/")
	////config.LoadYmlProfile("./resources/")
	//
	////fmt.Println(config.GetValueString("a.b"))
	////fmt.Println(config.GetValueBool("a.e"))
	////fmt.Println(config.GetValueIntDefault("a.f", 33))
	fmt.Println(config.GetValueObject("a.b"))

	////返回命令行参数后的其他参数
	//fmt.Println(flag.Args())
	////返回命令行参数后的其他参数个数
	//fmt.Println(flag.NArg())
	////返回使用的命令行参数个数
	//fmt.Println(flag.NFlag())

	//
	//if len(os.Args) < 0 {
	//	return
	//}
	//fmt.Printf("====%s====\n", os.Args[1])
	//fmt.Printf("====%s====\n", os.Args[2])
	//fmt.Printf("====%s====\n", os.Args[3])
	//fmt.Printf("%s\n", os.Args)

	//serviceLogger = log.GetLogger("service")
	//testLogger = log.GetLogger("test")
	//
	//r := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	////gin.DisableConsoleColor()
	//
	//r.Use(web.ResponseHandler())
	//r.GET("/get", get1)
	//r.GET("/test", test)
	//r.GET("/service", service)
	//
	//// 添加日志api到web
	//log.LogRouters(r)
	//
	//r.Run(":8082")
}

func get1(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    "12",
	})
}

func test(c *gin.Context) {
	testLogger.Debug("test-debug")
	testLogger.Info("test-debug")
	testLogger.Warn("test-debug")
	testLogger.Error("test-debug")
	//testLogger.Fatalf("test-debug")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "32",
		"message": "失败",
		"data":    "dfs",
	})
}

func service(c *gin.Context) {
	serviceLogger.Debug("service-debug")
	serviceLogger.Info("service-info")
	serviceLogger.Warn("service-warn")
	serviceLogger.Error("service-error")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    "12",
	})
}
