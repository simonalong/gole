# gole

gole 框架为基础工具包，只提供日志、配置、时间等基础工具，不提供与和二方框架、三方结合的工具封装，与二方框架、三方结合的工具类封装在gole-boot框架中

## 下载
```shell
go get github.com/simonalong/gole
```
### 提示：
1. 更新相关依赖
```shell
go mod tidy
```
2. 下载指定版本
```shell
go get github.com/simonalong/gole@<version>
```

## 快速入门
gole定位是基础工具框架，包含各种各样的基础工具

### 包列表
| 包名                    |          简介          |
|-----------------------|:--------------------:|
| [util](/util)         |      基础工具（更新中）       |
| [config](/config)     |        配置文件管理        |
| [validate](/validate) |         校验核查         |
| [logger](/logger)     |          日志          |
| [goid](/goid)         | 局部id传递处理（theadlocal） |
| [json](/jsonobj)         |     json字符串处理工具      |
| [time](/time)         |        时间管理工具        |
| [file](/file)         |        文件管理工具        |
| [coder](/coder)       |       编解码加解密工具       |
| [listener](/listener) |        事件监听机制        |
| [bean](/bean)         |        对象管理工具        |
| [excel](/excel)        |        excel包        |

### gole 测试
根目录提供go_test.sh文件，统一执行所有gole中包的测试模块
```shell
sh go_test.sh
```
