# GoleMap支持功能
GoleMap是连接golang实体和实体中的中间结构，是普通map的增强版本，提供如下功能
1. 提供类型转换：fromJson，fromMap，fromEntity，fromRows等功能，支持of初始化
2. 并发安全：使用currentMap保证并发安全
3. 提供有序性：原生map的输出是无序的，对一些功能不友好，这里支持有序
4. 提供与实体的转化功能：toEntity

### 初始化GoleMap
```go
// ---- 无序map ----
dataMap := maps.Of("a", 1, "b", 12, "k3", 13, "k4", 14)
dataMap := maps.New()

// ---- 有序map ----
dataMap := maps.OfSort("a", 1, "b", 12, "k3", 13, "k4", 14)
dataMap := maps.NewSort()

// xx为这三类：entity、map、string（json格式）; （有序map）
maps.From(xx)
// 也有单独的函数（有序map）
maps.FromJson(xx)
maps.FromMap(xx)
maps.FromEntity(xx)

// 也支持从td读取的内容格式转换；（有序map） 
FromRows(rows driver.Rows) []*GoleMap{}
```
### 添加值
```go
dataMap.Put("c", 124)
```
### 获取值
这里支持所有基本类型，也包括`time.Time`类型，也支持实体类型，不过实体类型为`ToEntity`
```go
Get(key string) (interface{}, bool)
GetInt(key string) (int, bool)
GetInt8(key string) (int8, bool)
GetInt16(key string) (int16, bool)
GetInt32(key string) (int32, bool)
GetInt64(key string) (int64, bool)
GetUInt8(key string) (uint8, bool)
GetUInt16(key string) (uint16, bool)
GetUInt32(key string) (uint32, bool)
GetUInt64(key string) (uint64, bool)
GetFloat32(key string) (float32, bool)
GetFloat64(key string) (float64, bool)
GetBool(key string) (bool, bool)
GetComplex64(key string) (complex64, bool)
GetComplex128(key string) (complex128, bool)
GetString(key string) (string, bool)
GetTime(key string) (time.Time, bool)
GetBytes(key string) ([]bytes, bool)
```
```go
// 转换为普通map
ToMap() map[string]interface{}
// 这个ToJson输出的不是无序的，因为Json这个不好搞成有序
ToJson() string {}
// 这个是有序的展示
ToString() string {}
// 入参为某实体的指针
ToEntity(pEntity interface{}) error {}
```
ToEntity 示例：
```go
type DemoEntity struct {
    Ts      time.Time `column:"ts"`
    Name    string    `column:"name"`
    Age     int       `column:"age"`
    Address string    `column:"address"`
}

entity1Expect := DemoEntity{}
err := entityMap.ToEntity(&entity1Expect)
```

### 循环
请使用如下循环，用来在有序情况下保证有序输出
```go
dataMap := maps.NewSort()
// ...
for _, key := range dataMap.Keys() {
    // ...
}
```

### 与实体转换
支持实体，其中的转换是将列名或者标签配置的别名转换为map，其中支持三类
1. column标签
2. json标签
3. 无标签：首字母变小写

优先级：column标签 > json标签 > 无标签

```go
type DemoEntity struct {
    Ts1      time.Time `column:"ts"`
    Name1    string    `column:"name"`
    Age1     int       `column:"age"`
    Address1 string    `column:"address"`
}
// 对应
["ts":"2024-07-17 11:03:24.225","name":"test","age":"22","address":"浙江"]
```
```go
type DemoEntity struct {
    Ts1      time.Time `json:"ts"`
    Name1    string    `json:"name"`
    Age1     int       `json:"age"`
    Address1 string    `json:"address"`
}
// 对应
["ts":"2024-07-17 11:03:24.225","name":"test","age":"22","address":"浙江"]
```
```go
type DemoEntity struct {
    Ts      time.Time
    Name    string
    Age     int
    Address string
}
// 对应
["ts":"2024-07-17 11:03:24.225","name":"test","age":"22","address":"浙江"]
```
