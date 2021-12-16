# tools
tools是go代码写的的个人工具包。

目前有以下功能
- yaml、properties、json、map互转的工具
- http的简单封装工具  
- 日志的工具封装
- web返回值的异常打印

### 1. yaml功能
提供如下格式的转换
```text
 1.yaml <---> properties
 2.yaml <---> json
 3.yaml <---> map
 4.yaml <---> list
 5.yaml <---> kvList
```
### 2. http 功能
提供http客户端的协议工具，对返回值增加结构的解析
```json
{
  "code": "xxx",
  "message": "xxx",
  "data": "xxx"
}
```

### 3. log 功能
1. 支持日志文件切分
2. 支持日志颜色
3. 增加logger维度
4. 增加对logger的日志级别管控

#### 日志路径配置用法

```go
// 文件: xxx.go

var appLog *logrus.Logger

func init() {
    // 在路径/home/isc-xxx-service/logs/路径下生成文件，app-info.log、app-warn.log、app-err.log，以及相关的切片日志，默认保存30天
    log.GetLoggerWithConfig("appLog", "/home/isc-xxx-service/logs/app", "/api/xxx/", true)
}
```

#### 日志管控
如下，添加日志的Router，添加后就可以对日志进行管控了
```go
func main() {
    engine := gin.Default()

    // 添加日志的管控api
    log.LogRouters(engine)

    engine.Run(":8082")
}
```
使用日志管控时候，调用如下，可以查看到可以管控的logger的命令，其中/api/xxx/是上面配置的
> curl http://localhost:port/api/xxx/help

```json
{
  "修改：host和port-----":"curl -X POST http://localhost:port/api/xxx/host/change/{host}/{port}",
  "修改：logger的级别----":"curl -X POST http://localhost:port/api/xxx/logger/level/{loggerName}/{level}",
  "修改：所有logger的级别":"curl -X POST http://localhost:port/api/xxx/logger/root/level/{level}",
  "查询：Logger集合-----":"curl http://localhost:port/api/xxx/logger/list",
  "查询：帮助-----------":"curl http://localhost:port/api/xxx/help"
}
```

### 4. 返回值异常打印
```go
func main() {
    r := gin.Default()

    // 配置：返回值异常情况的打印
    engine.Use(web.ResponseHandler())

    r.Run(":8082")
}
```
