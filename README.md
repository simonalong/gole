# gole
gole是go代码写的的个人工具包。

### 安装使用
```shell
go get github.com/simonalong/gole
```

目前有以下功能
- yaml、properties、json、map互转的工具
- http的简单封装工具
- config配置文件处理工具  
- 日志的工具封装
- web返回值的异常打印
- interface转换具体类型工具
- json转object工具：省去`json:"xxx"`
- message、cache  
- 其他的一些简单工具

## 1. yaml功能
提供如下格式的转换
```text
 1.yaml <---> properties
 2.yaml <---> json
 3.yaml <---> map
 4.yaml <---> list
 5.yaml <---> kvList
```
## 2. http 功能
提供http客户端的协议工具，对返回值增加结构的解析
```json
{
  "code": "xxx",
  "message": "xxx",
  "data": "xxx"
}
```

## 3. 配置文件 功能
给go项目提供配置文件加载规范，所有文件都放在 ./resource目录下，借鉴spring的规范，使用application-{profile}.mmm进行解析，支持yaml、yml、property、json等格式
### 1. 支持profile
`--gole.profile xxx`
即可读取./resource/application-xxx.mmm文件内容。获取配置内容可以使用config包的api获取即可

### a. 基本类型
```go
config.GetValueString(key string) string
config.GetValueInt(key string) int
config.GetValueInt8(key string) int8
config.GetValueInt16(key string) int16
config.GetValueInt32(key string) int32
config.GetValueInt64(key string) int64
config.GetValueBool(key string) bool
// ... 等等基本类型 ...
```

### b. 复杂类型
```go
// 也支持复杂类型，直接根据key也可以读取获得对应对象，使用如下方法，然后在使用util.DataToObject(data, obj)即可得到想要的结构
config.GetValueObject(key string) interface{}
```
示例：
```go
// 结构
type Entity struct {
    Name string
    Age int
}
```

```yaml
# ./resource/application.yml 配置文件
entity:
  name: "chen"
  age: 12
```
```go
// 获取
entity := Entity{}
config.GetValueObject("entity", &entity)
```

### c. 动态修改配置
这个功能是基于提供的在线修改功能直接修改，该功能会直接覆盖项目中的配置
```shell
curl -X POST http://localhost:port/api/gole/env -d '{"key":xxx, "value":xxx}'
```


## 3. log 功能
1. 支持日志文件切分
2. 支持日志颜色
3. 增加logger维度
4. 增加对logger的日志级别管控
5. 动态修改日志级别

### 日志路径配置用法

```go
// 文件: xxx.go

var appLog *logrus.Logger

func init() {
    // 在路径/home/isc-xxx-service/logs/路径下生成文件，app-info.log、app-warn.log、app-err.log，以及相关的切片日志，默认保存30天
    log.GetLoggerWithConfig("appLog", "/home/isc-xxx-service/logs/app", "/api/xxx/", true)
}
```

### 日志管控
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
  "修改：host和port-----":"curl -X POST http://localhost:port/api/log/host/change/{host}/{port}",
  "修改：logger的级别----":"curl -X POST http://localhost:port/api/log/logger/level/{loggerName}/{level}",
  "修改：所有logger的级别":"curl -X POST http://localhost:port/api/log/logger/root/level/{level}",
  "修改：环境变量--------":"curl -X POST http://localhost:port/api/log/env",
  "查询：Logger集合-----":"curl http://localhost:port/api/log/logger/list",
  "查询：帮助-----------":"curl http://localhost:port/api/log/help"
}
```

## 4. gin框架返回值异常打印
这里使用的是gin框架
```go
func main() {
    r := gin.Default()

    // 配置：返回值异常情况的打印
    engine.Use(web.ResponseHandler())

    r.Run(":8082")
}
```
## 5. interface转换具体类型工具
在使用到interface时候，我们有时候需要使用具体类型，这种在go里面比较烦，就做了个简单的工具
```go
// ToInt ：  interface{} --> int
// ToInt8 ： interface{} --> int8
// ToInt16 ：interface{} --> int16
// ToInt32 ：interface{} --> int32
// ToInt64 ：interface{} --> int64

// ToUInt ：  interface{} --> uint
// ToUInt8 ： interface{} --> uint8
// ToUInt16 ：interface{} --> uint16
// ToUInt32 ：interface{} --> uint32
// ToUInt64 ：interface{} --> uint64

// ToFloat32 ：interface{} --> float32
// ToFloat64 ：interface{} --> float64

// ToBool ：interface{} --> bool

// ToComplex64 ：interface{} --> complex64
// ToComplex128 ：interface{} --> complex128

// ToMap ：interface{} --> map[string]interface{}

// Cast：向指定类型转换，比如：string转int
```
## 6. json和object互转工具
这个工具是为了解决json反解析到对象这种时候，类型的tag中必须使用`json:"xxx"`这种，代码对外的实体中会存在非常多的这种，如下
比如：
```go
type AppManagerUpdateReq struct {
    Id           int64  `json:"id"`
    AppName      string `json:"appName"`
    AppDesc      string `json:"appDesc"`
    ActiveStatus int8   `json:"activeStatus"`
}
```
为了去掉后面的json，这里做了反解析化工具。提供如下api进行转换
```go
// MapToObject              map         ——> 对象
// ArrayToObject            array       ——> 对象
// StrToObject              字符         ——> 对象
// ReaderJsonToObject       io.Reader   ——> 对象
// DataToObject：通用转换     以上类型     ——> 对象
//
// ObjectToJson             对象         ——>json字符：对应的json中的字段为小写
// ObjectToData：通用转换     对象         ——>转换后的对象
```
#### 提示：
这里的转换支持如下三种特性
- 无json标示：反解析不需要添加'json:"xxx"'，而且反解析时候对应的map中的key大小写均可
- 类型无限制：map中对应的类型只要能转换进去即可，比如：map中的value值为string类型，但是存储的是数字，实体为int，也是可以的
- 自定义类型：比如自定义类型`type myType int8` 这种类型修饰的也可以转换
- 不支持：暂时不支持属性为指针的类型转换

### 举例
#### 1. 无`json:"xxx"`解析：json到对象
```go
type ValueInnerEntityT struct {
    AppName string
    Age  int
}

func TestMapToObjectT(t *testing.T) {
    inner1 := map[string]interface{}{}
    inner1["appName"] = "inner_1"
    inner1["age"] = 1

    var targetObj ValueInnerEntityT
    util.MapToObject(inner1, &targetObj)
    Equal(t, util.ToJsonString(targetObj), "{\"AppName\":\"inner_1\",\"Age\":1}")
}
```
嵌套结构也支持
```go
type ValueInnerEntity1 struct {
    Name string
    Age  int
}

type ValueInnerEntity2 struct {
    Name   string
    Age    int
    Inner1 ValueInnerEntity1
}

func TestMapToObject2(t *testing.T) {
    inner1 := map[string]interface{}{}
    inner1["name"] = "inner_1"
    inner1["age"] = 1

    inner2 := map[string]interface{}{}
    inner2["name"] = "inner_2"
    inner2["age"] = 2
    inner2["inner1"] = inner1

    var targetObj ValueInnerEntity2
    util.MapToObject(inner2, &targetObj)
    Equal(t, "{\"Name\":\"inner_2\",\"Age\":2,\"Inner1\":{\"Name\":\"inner_1\",\"Age\":1}}", util.ToJsonString(targetObj))
}
```
#### 2. 无`json:"xxx"`解析：对象到json
```go
type ValueObjectTest1 struct {
    AppName string
    Age  int
}

func TestObjectToJson1(t *testing.T) {
    entity := ValueObjectTest1{AppName: "zhou", Age: 12}
    Equal(t, util.ObjectToJson(entity), "{\"age\":12,\"appName\":\"zhou\"}")
}
```
结果
```json
{
    "age":12,
    "appName":"zhou"
}
```
复杂结构
```go
type ValueObjectTest3 struct {
    AppName []string
    Age1  map[string]interface{}
}

type ValueObjectTest4 struct {
    AppName string
    Inner  ValueObjectTest3
}

func TestObjectToJson4(t *testing.T) {
    var arrays []string
    arrays = append(arrays, "zhou")
    arrays = append(arrays, "wang")

    dataMap := map[string]interface{}{}
    dataMap["a"] = 1
    dataMap["b"] = 2

    entity3 := ValueObjectTest3 {
        AppName: arrays,
        Age1: dataMap,
    }

    var entity4 ValueObjectTest4
    entity4.Inner = entity3
    entity4.AppName = "zhou"
    Equal(t, util.ObjectToJson(entity4), "{\"appName\":\"zhou\",\"inner\":{\"age1\":{\"a\":1,\"b\":2},\"appName\":[\"zhou\",\"wang\"]}}")
}
```
结果
```json
{
    "appName":"zhou",
    "inner":{
        "age1":{
            "a":1,
            "b":2
        },
        "appName":[
            "zhou",
            "wang"
        ]
    }
}
```
#### 3. 类型兼容（互转）
如下string类型也可以转换为int
```go
type ValueInnerEntityTem struct {
    Name string
    Age  int
}

func TestMapToObjectTem(t *testing.T) {
    inner1 := map[string]interface{}{}
    inner1["name"] = "inner_1"
    inner1["age"] = "123"

    var targetObj ValueInnerEntity1
    _ = util.MapToObject(inner1, &targetObj)
    Equal(t, "{\"Name\":\"inner_1\",\"Age\":123}", util.ToJsonString(targetObj))
}
```
#### 4. 支持自定义类型
```go
type MyEnum int

type ValueInnerEntityTem struct {
    Name string
    Age  MyEnum
}

func TestMapToObjectTem(t *testing.T) {
    inner1 := map[string]interface{}{}
    inner1["name"] = "inner_1"
    inner1["age"] = "1"

    var targetObj ValueInnerEntity1
    _ = util.MapToObject(inner1, &targetObj)
    Equal(t, util.ToJsonString(targetObj), "{\"Name\":\"inner_1\",\"Age\":1}")
}
```
