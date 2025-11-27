# jsonobj

jsonobj包是为了方便对json格式的字符处理使用；底层是封装了jsonpath包（github.com/oliveagle/jsonpath），支持jsonpath的操作

```json
{
    "k1":12,
    "k2":true,
    "k3":{
        "k31":32,
        "k32":"str",
        "k33":{
            "k331":12
        }
    }
}
```

```go
func TestLoad(t *testing.T) {
    // str就是上面的字符串
    jsonObject, err := jsonobj.Load(str)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    assert.Equal(t, jsonObject.Get("$.k1"), 12)
    assert.Equal(t, jsonObject.Get("$.k2"), true)
    // 支持连key获取
    assert.Equal(t, jsonObject.Get("$.k3.k31"), 32)
    assert.Equal(t, jsonObject.Get("$.k3.k32"), "str")
    assert.Equal(t, jsonObject.Get("$.k3.k33.k331"), 12)
}
```

另外也提供了各种类型的api用于指定的类型使用
```go

GetString(key string) string
GetInt(key string) int
GetInt8(key string) int8
GetInt16(key string) int16
GetInt32(key string) int32
GetInt64(key string) int64
GetUInt(key string) uint
GetUInt8(key string) uint8
GetUInt16(key string) uint16
GetUInt32(key string) uint32
GetUInt64(key string) uint64
GetFloat32(key string) float32
GetFloat64(key string) float64
GetBool(key string) bool
GetObject(key string, targetPtrObj any) error
GetArray(key string) []any
```

如下全都是框架（github.com/oliveagle/jsonpath）支持的，更多具体功能请见该框架

| 操作                        |  支持  | 描述                     |
|---------------------------|:----:|------------------------|
| $ 					     |  Y   | 要查询的根元素。这将启动所有路径表达式。   |
| @ 				         |  Y   | 当前节点正由筛选器谓词处理          |
| * 					     |  X   | 通配符。任何地方都可以使用，需要名称或数字 |
| .. 					     |  X   | 深度扫描。任何需要名称的地方都可以使用  |
| .<name> 				     |  Y   | 点号的次级                  |
| ['<name>' (, '<name>')]   |  X   | 括号标记的次级                |
| [<number> (, <number>)]   |  Y   | 数组索引                   |
| [start:end] 			     |  Y   | 数组切片运算符                |
| [?(<expression>)] 	     |  Y   | 筛选器表达式。表达式的计算结果必须为布尔值。 |


| jsonpath | result|
| :--------- | :-------|
| $.expensive 			                           | 10|
| $.store.book[0].price                            | 8.95|
| $.store.book[-1].isbn                            | "0-395-19395-8"|
| $.store.book[0,1].price                          | [8.95, 12.99]   |
| $.store.book[0:2].price                          | [8.95, 12.99, 8.99]|
| $.store.book[?(@.isbn)].price                    |  [8.99, 22.99] |
| $.store.book[?(@.price > 10)].title              | ["Sword of Honour", "The Lord of the Rings"]|
| $.store.book[?(@.price < $.expensive)].price     | [8.95, 8.99] |
| $.store.book[:].price                            | [8.9.5, 12.99, 8.9.9, 22.99] |
| $.store.book[?(@.author =~ /(?i).*REES/)].author | "Nigel Rees" |
