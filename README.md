## tools
tools是go代码写的的个人工具包。

目前有以下功能
- yaml、properties、json、map互转的工具
- http的简单封装工具  
- 日志的工具封装： 
  - 文件自动切分级别划分
  - 日志配置颜色
  - 控制日志logger的级别  


### yaml用法
### http用法
### log 工具

```go
func main() {
	// 配置日志路径，会在对应目录下生成文件，biz-debug.log、biz-warn.log、biz-error.log、biz-fatal.log
	log.LogPathSet("/user/xxxx/logs/biz")
	// 日志管理api的前缀
	log.LogApiConfig("/api/core/troy")
	// 是否配置日志颜色
	log.LogColor(true)

	// 获取对应的logger
	bizLogger = log.GetLogger("biz")

	r := gin.Default()

	// 添加日志api到web
	log.LogRouters(r)

	r.Run(":8082")
}
```
