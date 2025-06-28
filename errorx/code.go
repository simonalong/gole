package errorx

import "github.com/simonalong/gole/util"

var (
	// ------------------- 10x -------------------

	// SC_CONTINUE 100
	SC_CONTINUE *BaseError
	// SC_SWITCHING_PROTOCOLS 101
	SC_SWITCHING_PROTOCOLS *BaseError
	// SC_PROCESSING 102
	SC_PROCESSING *BaseError

	// ------------------- 20x -------------------

	// SC_OK 200
	SC_OK                            *BaseError
	SC_CREATED                       *BaseError
	SC_ACCEPTED                      *BaseError
	SC_NON_AUTHORITATIVE_INFORMATION *BaseError
	SC_NO_CONTENT                    *BaseError
	SC_RESET_CONTENT                 *BaseError
	SC_PARTIAL_CONTENT               *BaseError

	// ------------------- 30x -------------------

	SC_MULTIPLE_CHOICES   *BaseError
	SC_MOVED_PERMANENTLY  *BaseError
	SC_FOUND              *BaseError
	SC_SEE_OTHER          *BaseError
	SC_NOT_MODIFIED       *BaseError
	SC_USE_PROXY          *BaseError
	SC_TEMPORARY_REDIRECT *BaseError

	// ------------------- 40x -------------------

	SC_BAD_REQUEST                     *BaseError
	SC_UNAUTHORIZED                    *BaseError
	SC_PAYMENT_REQUIRED                *BaseError
	SC_FORBIDDEN                       *BaseError
	SC_NOT_FOUND                       *BaseError
	SC_METHOD_NOT_ALLOWED              *BaseError
	SC_NOT_ACCEPTABLE                  *BaseError
	SC_PROXY_AUTHENTICATION_REQUIRED   *BaseError
	SC_REQUEST_TIMEOUT                 *BaseError
	SC_CONFLICT                        *BaseError
	SC_GONE                            *BaseError
	SC_LENGTH_REQUIRED                 *BaseError
	SC_PRECONDITION_FAILED             *BaseError
	SC_REQUEST_TOO_LONG                *BaseError
	SC_REQUEST_URI_TOO_LONG            *BaseError
	SC_UNSUPPORTED_MEDIA_TYPE          *BaseError
	SC_REQUESTED_RANGE_NOT_SATISFIABLE *BaseError
	SC_EXPECTATION_FAILED              *BaseError
	SC_UNPROCESSABLE_ENTITY            *BaseError
	SC_LOCKED                          *BaseError
	SC_FAILED_DEPENDENCY               *BaseError
	SC_OUTOFRANGE                      *BaseError

	// ------------------- 50x -------------------

	SC_SERVER_ERROR               *BaseError
	SC_NOT_IMPLEMENTED            *BaseError
	SC_BAD_GATEWAY                *BaseError
	SC_SERVICE_UNAVAILABLE        *BaseError
	SC_GATEWAY_TIMEOUT            *BaseError
	SC_HTTP_VERSION_NOT_SUPPORTED *BaseError
	SC_INSUFFICIENT_STORAGE       *BaseError
	SC_UNKNOWN_ERR                *BaseError
	SC_NOTFOUND_ERR               *BaseError
	SC_ALREADY_EXISTS_ERR         *BaseError
	SC_RESOURCE_EXHAUSTED_ERR     *BaseError
	SC_ABORTED_ERR                *BaseError
	SC_DATALOSS_ERR               *BaseError

	// ------------------- 60x -------------------

	SC_THIRD_ERR    *BaseError
	SC_HTTP_ERR     *BaseError
	SC_TCP_ERR      *BaseError
	SC_UDP_ERR      *BaseError
	SC_RPC_ERR      *BaseError
	SC_GRPC_ERR     *BaseError
	SC_DB_ERR       *BaseError
	SC_MQ_ERR       *BaseError
	SC_REDIS_ERR    *BaseError
	SC_TDENGINE_ERR *BaseError
	SC_FILE_ERR     *BaseError
	SC_MYSQL_ERR    *BaseError
	SC_NATS_ERR     *BaseError
)

var ScCodeMap map[string]*ScCode
var ScCodeErrMap map[string]*BaseError

// GrpcStatusMap grpc状态码和标准错误的对应关系
var GrpcStatusMap map[uint32]*BaseError
var ScCodeOkMap util.ISCSet[string]

const DefaultScHttpCode = 500
const DefaultScGrpcCode = 13

type ScCode struct {
	HttpCode int
	GrpcCode uint32
}

func init() {
	initCodeMsg()
	initCodeMap()
	initCodeErrMap()
	initCodeOkMap()
	initGrpcStatusMap()
}

func initCodeMsg() {
	// SC_CONTINUE 100
	SC_CONTINUE = New("SC_CONTINUE", "客户端继续")
	// SC_CONTINUE 101
	SC_SWITCHING_PROTOCOLS = New("SC_SWITCHING_PROTOCOLS", "服务器正在根据Upgrade报头切换协议")
	// SC_CONTINUE 102
	SC_PROCESSING = New("SC_PROCESSING", "交换协议")

	// ------------------- 20x -------------------

	// ScOk 200
	SC_OK = New("SC_OK", "成功")
	//
	SC_CREATED = New("SC_CREATED", "请求创建资源成功")
	SC_ACCEPTED = New("SC_ACCEPTED", "请求被接受")
	SC_NON_AUTHORITATIVE_INFORMATION = New("SC_NON_AUTHORITATIVE_INFORMATION", "客户端提供的元信息并非来自服务器")
	SC_NO_CONTENT = New("SC_NO_CONTENT", "请求成功无信息返回")
	SC_RESET_CONTENT = New("SC_RESET_CONTENT", "请求代理重置")
	SC_PARTIAL_CONTENT = New("SC_PARTIAL_CONTENT", "完成部分GET请求")

	// ------------------- 30x -------------------

	SC_MULTIPLE_CHOICES = New("SC_MULTIPLE_CHOICES", "多选择")
	SC_MOVED_PERMANENTLY = New("SC_MOVED_PERMANENTLY", "资源已迁移，请使用新的url")
	SC_FOUND = New("SC_FOUND", "资源已临时迁移，旧url仍可用")
	SC_SEE_OTHER = New("SC_SEE_OTHER", "查看其他url响应的数据")
	SC_NOT_MODIFIED = New("SC_NOT_MODIFIED", "GET操作发现资源可用且未被修改")
	SC_USE_PROXY = New("SC_USE_PROXY", "请使用代理")
	SC_TEMPORARY_REDIRECT = New("SC_TEMPORARY_REDIRECT", "临时重定向")

	// ------------------- 40x -------------------

	SC_BAD_REQUEST = New("SC_BAD_REQUEST", "请求异常")
	SC_UNAUTHORIZED = New("SC_UNAUTHORIZED", "需要HTTP身份验证")
	SC_PAYMENT_REQUIRED = New("SC_PAYMENT_REQUIRED", "保留以备将来使用")
	SC_FORBIDDEN = New("SC_FORBIDDEN", "无权限")
	SC_NOT_FOUND = New("SC_NOT_FOUND", "请求资源不可用")
	SC_METHOD_NOT_ALLOWED = New("SC_METHOD_NOT_ALLOWED", "Request-URI方法禁止访问")
	SC_NOT_ACCEPTABLE = New("SC_NOT_ACCEPTABLE", "不接受")
	SC_PROXY_AUTHENTICATION_REQUIRED = New("SC_PROXY_AUTHENTICATION_REQUIRED", "代理客户端必须首先通过进行身份验证")
	SC_REQUEST_TIMEOUT = New("SC_REQUEST_TIMEOUT", "请求超时")
	SC_CONFLICT = New("SC_CONFLICT", "资源冲突，请求失败")
	SC_GONE = New("SC_GONE", "资源在服务器上不再可用")
	SC_LENGTH_REQUIRED = New("SC_LENGTH_REQUIRED", "未定义Content-Length，无法处理请求")
	SC_PRECONDITION_FAILED = New("SC_PRECONDITION_FAILED", "前置条件评估失败")
	SC_REQUEST_TOO_LONG = New("SC_REQUEST_TOO_LONG", "请求实体过长")
	SC_REQUEST_URI_TOO_LONG = New("SC_REQUEST_URI_TOO_LONG", "请求Url过长")
	SC_UNSUPPORTED_MEDIA_TYPE = New("SC_UNSUPPORTED_MEDIA_TYPE", "请求的实体的格式不支持")
	SC_REQUESTED_RANGE_NOT_SATISFIABLE = New("SC_REQUESTED_RANGE_NOT_SATISFIABLE", "请求的字节范围不支持")
	SC_EXPECTATION_FAILED = New("SC_EXPECTATION_FAILED", "请求标头不支持")
	//SC_INSUFFICIENT_SPACE_ON_RESOURCE = New("SC_INSUFFICIENT_SPACE_ON_RESOURCE", "客户端继续")
	//SC_METHOD_FAILURE = New("SC_METHOD_FAILURE", "客户端继续")
	SC_UNPROCESSABLE_ENTITY = New("SC_UNPROCESSABLE_ENTITY", "不可处理实体")
	SC_LOCKED = New("SC_LOCKED", "锁定")
	SC_FAILED_DEPENDENCY = New("SC_FAILED_DEPENDENCY", "依赖失败")
	SC_OUTOFRANGE = New("SC_OUTOFRANGE", "越界")

	// ------------------- 50x -------------------

	SC_SERVER_ERROR = New("SC_SERVER_ERROR", "服务端异常")
	SC_NOT_IMPLEMENTED = New("SC_NOT_IMPLEMENTED", "未实现")
	SC_BAD_GATEWAY = New("SC_BAD_GATEWAY", "错误网关")
	SC_SERVICE_UNAVAILABLE = New("SC_SERVICE_UNAVAILABLE", "服务不可用")
	SC_GATEWAY_TIMEOUT = New("SC_GATEWAY_TIMEOUT", "网关超时")
	SC_HTTP_VERSION_NOT_SUPPORTED = New("SC_HTTP_VERSION_NOT_SUPPORTED", "不支持的HTTP版本")
	SC_INSUFFICIENT_STORAGE = New("SC_INSUFFICIENT_STORAGE", "存储空间不足")
	SC_UNKNOWN_ERR = New("SC_UNKNOWN_ERR", "未知异常")
	SC_NOTFOUND_ERR = New("SC_NOTFOUND_ERR", "找不到指定资源")
	SC_ALREADY_EXISTS_ERR = New("SC_ALREADY_EXISTS_ERR", "已经存在")
	SC_RESOURCE_EXHAUSTED_ERR = New("SC_RESOURCE_EXHAUSTED_ERR", "资源耗尽")
	SC_ABORTED_ERR = New("SC_ABORTED_ERR", "操作终止")
	SC_DATALOSS_ERR = New("SC_DATALOSS_ERR", "数据丢失")

	// ------------------- 60x -------------------

	SC_THIRD_ERR = New("SC_THIRD_ERR", "调用三方异常")
	SC_HTTP_ERR = New("SC_HTTP_ERR", "调用http异常")
	SC_TCP_ERR = New("SC_TCP_ERR", "调用tcp异常")
	SC_UDP_ERR = New("SC_UDP_ERR", "调用udp异常")
	SC_RPC_ERR = New("SC_RPC_ERR", "调用rpc异常")
	SC_GRPC_ERR = New("SC_GRPC_ERR", "调用grpc异常")
	SC_DB_ERR = New("SC_DB_ERR", "调用数据库异常")
	SC_MQ_ERR = New("SC_MQ_ERR", "调用mq异常")
	SC_REDIS_ERR = New("SC_REDIS_ERR", "调用redis异常")
	SC_TDENGINE_ERR = New("SC_TDENGINE_ERR", "调用tdengine异常")
	SC_FILE_ERR = New("SC_FILE_ERR", "调用文件异常")
	SC_MYSQL_ERR = New("SC_MYSQL_ERR", "调用mysql异常")
	SC_NATS_ERR = New("SC_NATS_ERR", "调用nats异常")
}

func initCodeOkMap() {
	ScCodeOkMap = util.NewSet[string]()
	_ = ScCodeOkMap.Add("SC_CONTINUE")
	_ = ScCodeOkMap.Add("SC_SWITCHING_PROTOCOLS")
	_ = ScCodeOkMap.Add("SC_PROCESSING")

	// ------------------- 20x -------------------

	_ = ScCodeOkMap.Add("SC_OK")
	_ = ScCodeOkMap.Add("SC_CREATED")
	_ = ScCodeOkMap.Add("SC_ACCEPTED")
	_ = ScCodeOkMap.Add("SC_NON_AUTHORITATIVE_INFORMATION")
	_ = ScCodeOkMap.Add("SC_NO_CONTENT")
	_ = ScCodeOkMap.Add("SC_RESET_CONTENT")
	_ = ScCodeOkMap.Add("SC_PARTIAL_CONTENT")
}

func initCodeErrMap() {
	ScCodeErrMap = make(map[string]*BaseError)
	ScCodeErrMap["SC_CONTINUE"] = SC_CONTINUE
	ScCodeErrMap["SC_SWITCHING_PROTOCOLS"] = SC_SWITCHING_PROTOCOLS
	ScCodeErrMap["SC_PROCESSING"] = SC_PROCESSING
	// ------------------- 20x -------------------
	ScCodeErrMap["SC_OK"] = SC_OK
	ScCodeErrMap["SC_CREATED"] = SC_CREATED
	ScCodeErrMap["SC_ACCEPTED"] = SC_ACCEPTED
	ScCodeErrMap["SC_NON_AUTHORITATIVE_INFORMATION"] = SC_NON_AUTHORITATIVE_INFORMATION
	ScCodeErrMap["SC_NO_CONTENT"] = SC_NO_CONTENT
	ScCodeErrMap["SC_RESET_CONTENT"] = SC_RESET_CONTENT
	ScCodeErrMap["SC_PARTIAL_CONTENT"] = SC_PARTIAL_CONTENT
	// ------------------- 30x -------------------
	ScCodeErrMap["SC_MULTIPLE_CHOICES"] = SC_MULTIPLE_CHOICES
	ScCodeErrMap["SC_MOVED_PERMANENTLY"] = SC_MOVED_PERMANENTLY
	ScCodeErrMap["SC_FOUND"] = SC_FOUND
	ScCodeErrMap["SC_SEE_OTHER"] = SC_SEE_OTHER
	ScCodeErrMap["SC_NOT_MODIFIED"] = SC_NOT_MODIFIED
	ScCodeErrMap["SC_USE_PROXY"] = SC_USE_PROXY
	ScCodeErrMap["SC_TEMPORARY_REDIRECT"] = SC_TEMPORARY_REDIRECT
	// ------------------- 40x -------------------
	ScCodeErrMap["SC_BAD_REQUEST"] = SC_BAD_REQUEST
	ScCodeErrMap["SC_UNAUTHORIZED"] = SC_UNAUTHORIZED
	ScCodeErrMap["SC_PAYMENT_REQUIRED"] = SC_PAYMENT_REQUIRED
	ScCodeErrMap["SC_FORBIDDEN"] = SC_FORBIDDEN
	ScCodeErrMap["SC_NOT_FOUND"] = SC_NOT_FOUND
	ScCodeErrMap["SC_METHOD_NOT_ALLOWED"] = SC_METHOD_NOT_ALLOWED
	ScCodeErrMap["SC_NOT_ACCEPTABLE"] = SC_NOT_ACCEPTABLE
	ScCodeErrMap["SC_PROXY_AUTHENTICATION_REQUIRED"] = SC_PROXY_AUTHENTICATION_REQUIRED
	ScCodeErrMap["SC_REQUEST_TIMEOUT"] = SC_REQUEST_TIMEOUT
	ScCodeErrMap["SC_CONFLICT"] = SC_CONFLICT
	ScCodeErrMap["SC_GONE"] = SC_GONE
	ScCodeErrMap["SC_LENGTH_REQUIRED"] = SC_LENGTH_REQUIRED
	ScCodeErrMap["SC_PRECONDITION_FAILED"] = SC_PRECONDITION_FAILED
	ScCodeErrMap["SC_REQUEST_TOO_LONG"] = SC_REQUEST_TOO_LONG
	ScCodeErrMap["SC_REQUEST_URI_TOO_LONG"] = SC_REQUEST_URI_TOO_LONG
	ScCodeErrMap["SC_UNSUPPORTED_MEDIA_TYPE"] = SC_UNSUPPORTED_MEDIA_TYPE
	ScCodeErrMap["SC_REQUESTED_RANGE_NOT_SATISFIABLE"] = SC_REQUESTED_RANGE_NOT_SATISFIABLE
	ScCodeErrMap["SC_EXPECTATION_FAILED"] = SC_EXPECTATION_FAILED
	ScCodeErrMap["SC_UNPROCESSABLE_ENTITY"] = SC_UNPROCESSABLE_ENTITY
	ScCodeErrMap["SC_LOCKED"] = SC_LOCKED
	ScCodeErrMap["SC_FAILED_DEPENDENCY"] = SC_FAILED_DEPENDENCY
	ScCodeErrMap["SC_OUTOFRANGE"] = SC_OUTOFRANGE
	// ------------------- 50x -------------------
	ScCodeErrMap["SC_SERVER_ERROR"] = SC_SERVER_ERROR
	ScCodeErrMap["SC_NOT_IMPLEMENTED"] = SC_NOT_IMPLEMENTED
	ScCodeErrMap["SC_BAD_GATEWAY"] = SC_BAD_GATEWAY
	ScCodeErrMap["SC_SERVICE_UNAVAILABLE"] = SC_SERVICE_UNAVAILABLE
	ScCodeErrMap["SC_GATEWAY_TIMEOUT"] = SC_GATEWAY_TIMEOUT
	ScCodeErrMap["SC_HTTP_VERSION_NOT_SUPPORTED"] = SC_HTTP_VERSION_NOT_SUPPORTED
	ScCodeErrMap["SC_INSUFFICIENT_STORAGE"] = SC_INSUFFICIENT_STORAGE
	ScCodeErrMap["SC_UNKNOWN_ERR"] = SC_UNKNOWN_ERR
	ScCodeErrMap["SC_NOTFOUND_ERR"] = SC_NOTFOUND_ERR
	ScCodeErrMap["SC_ALREADY_EXISTS_ERR"] = SC_ALREADY_EXISTS_ERR
	ScCodeErrMap["SC_RESOURCE_EXHAUSTED_ERR"] = SC_RESOURCE_EXHAUSTED_ERR
	ScCodeErrMap["SC_ABORTED_ERR"] = SC_ABORTED_ERR
	ScCodeErrMap["SC_DATALOSS_ERR"] = SC_DATALOSS_ERR
	// ------------------- 60x -------------------
	ScCodeErrMap["SC_THIRD_ERR"] = SC_THIRD_ERR
	ScCodeErrMap["SC_HTTP_ERR"] = SC_HTTP_ERR
	ScCodeErrMap["SC_TCP_ERR"] = SC_TCP_ERR
	ScCodeErrMap["SC_UDP_ERR"] = SC_UDP_ERR
	ScCodeErrMap["SC_RPC_ERR"] = SC_RPC_ERR
	ScCodeErrMap["SC_GRPC_ERR"] = SC_GRPC_ERR
	ScCodeErrMap["SC_DB_ERR"] = SC_DB_ERR
	ScCodeErrMap["SC_MQ_ERR"] = SC_MQ_ERR
	ScCodeErrMap["SC_REDIS_ERR"] = SC_REDIS_ERR
	ScCodeErrMap["SC_TDENGINE_ERR"] = SC_TDENGINE_ERR
	ScCodeErrMap["SC_FILE_ERR"] = SC_FILE_ERR
	ScCodeErrMap["SC_MYSQL_ERR"] = SC_MYSQL_ERR
	ScCodeErrMap["SC_NATS_ERR"] = SC_NATS_ERR
}

func initCodeMap() {
	ScCodeMap = make(map[string]*ScCode)
	ScCodeMap["SC_CONTINUE"] = NewScCodeEntity(100, 0)
	ScCodeMap["SC_SWITCHING_PROTOCOLS"] = NewScCodeEntity(101, 0)
	ScCodeMap["SC_PROCESSING"] = NewScCodeEntity(102, 0)

	// ------------------- 20x -------------------

	ScCodeMap["SC_OK"] = NewScCodeEntity(200, 0)
	ScCodeMap["SC_CREATED"] = NewScCodeEntity(201, 0)
	ScCodeMap["SC_ACCEPTED"] = NewScCodeEntity(202, 0)
	ScCodeMap["SC_NON_AUTHORITATIVE_INFORMATION"] = NewScCodeEntity(203, 0)
	ScCodeMap["SC_NO_CONTENT"] = NewScCodeEntity(204, 0)
	ScCodeMap["SC_RESET_CONTENT"] = NewScCodeEntity(205, 0)
	ScCodeMap["SC_PARTIAL_CONTENT"] = NewScCodeEntity(206, 0)

	// ------------------- 30x -------------------
	ScCodeMap["SC_MULTIPLE_CHOICES"] = NewScCodeEntity(300, 0)
	ScCodeMap["SC_MOVED_PERMANENTLY"] = NewScCodeEntity(301, 0)
	ScCodeMap["SC_MOVED_TEMPORARILY"] = NewScCodeEntity(302, 0)
	ScCodeMap["SC_SEE_OTHER"] = NewScCodeEntity(303, 0)
	ScCodeMap["SC_NOT_MODIFIED"] = NewScCodeEntity(304, 0)
	ScCodeMap["SC_USE_PROXY"] = NewScCodeEntity(305, 0)
	ScCodeMap["SC_TEMPORARY_REDIRECT"] = NewScCodeEntity(307, 0)

	// ------------------- 40x -------------------

	ScCodeMap["SC_BAD_REQUEST"] = NewScCodeEntity(400, 3)
	ScCodeMap["SC_UNAUTHORIZED"] = NewScCodeEntity(401, 16)
	ScCodeMap["SC_PAYMENT_REQUIRED"] = NewScCodeEntity(402, 3)
	ScCodeMap["SC_FORBIDDEN"] = NewScCodeEntity(403, 7)
	ScCodeMap["SC_NOT_FOUND"] = NewScCodeEntity(404, 3)
	ScCodeMap["SC_METHOD_NOT_ALLOWED"] = NewScCodeEntity(405, 9)
	ScCodeMap["SC_NOT_ACCEPTABLE"] = NewScCodeEntity(406, 3)
	ScCodeMap["SC_PROXY_AUTHENTICATION_REQUIRED"] = NewScCodeEntity(407, 3)
	ScCodeMap["SC_REQUEST_TIMEOUT"] = NewScCodeEntity(408, 4)
	ScCodeMap["SC_CONFLICT"] = NewScCodeEntity(409, 3)
	ScCodeMap["SC_GONE"] = NewScCodeEntity(410, 3)
	ScCodeMap["SC_LENGTH_REQUIRED"] = NewScCodeEntity(411, 3)
	ScCodeMap["SC_PRECONDITION_FAILED"] = NewScCodeEntity(412, 3)
	ScCodeMap["SC_REQUEST_TOO_LONG"] = NewScCodeEntity(413, 3)
	ScCodeMap["SC_REQUEST_URI_TOO_LONG"] = NewScCodeEntity(414, 3)
	ScCodeMap["SC_UNSUPPORTED_MEDIA_TYPE"] = NewScCodeEntity(415, 3)
	ScCodeMap["SC_REQUESTED_RANGE_NOT_SATISFIABLE"] = NewScCodeEntity(416, 3)
	ScCodeMap["SC_EXPECTATION_FAILED"] = NewScCodeEntity(417, 3)
	//ScCodeMap["SC_INSUFFICIENT_SPACE_ON_RESOURCE"] = NewScCodeEntity(100, 3)
	//ScCodeMap["SC_METHOD_FAILURE"] = NewScCodeEntity(100, 3)
	ScCodeMap["SC_UNPROCESSABLE_ENTITY"] = NewScCodeEntity(422, 3)
	ScCodeMap["SC_LOCKED"] = NewScCodeEntity(423, 3)
	ScCodeMap["SC_FAILED_DEPENDENCY"] = NewScCodeEntity(424, 3)
	ScCodeMap["SC_OUTOFRANGE"] = NewScCodeEntity(400, 11)

	// ------------------- 50x -------------------

	ScCodeMap["SC_SERVER_ERROR"] = NewScCodeEntity(500, 13)
	ScCodeMap["SC_NOT_IMPLEMENTED"] = NewScCodeEntity(501, 12)
	ScCodeMap["SC_BAD_GATEWAY"] = NewScCodeEntity(502, 13)
	ScCodeMap["SC_SERVICE_UNAVAILABLE"] = NewScCodeEntity(503, 14)
	ScCodeMap["SC_GATEWAY_TIMEOUT"] = NewScCodeEntity(504, 13)
	ScCodeMap["SC_HTTP_VERSION_NOT_SUPPORTED"] = NewScCodeEntity(505, 13)
	ScCodeMap["SC_INSUFFICIENT_STORAGE"] = NewScCodeEntity(507, 13)
	ScCodeMap["SC_UNKNOWN_ERR"] = NewScCodeEntity(500, 2)
	ScCodeMap["SC_NOTFOUND_ERR"] = NewScCodeEntity(500, 5)
	ScCodeMap["SC_ALREADY_EXISTS_ERR"] = NewScCodeEntity(500, 6)
	ScCodeMap["SC_RESOURCE_EXHAUSTED_ERR"] = NewScCodeEntity(500, 8)
	ScCodeMap["SC_ABORTED_ERR"] = NewScCodeEntity(500, 10)
	ScCodeMap["SC_DATALOSS_ERR"] = NewScCodeEntity(500, 15)

	// ------------------- 60x -------------------

	ScCodeMap["SC_THIRD_ERR"] = NewScCodeEntity(600, 13)
	ScCodeMap["SC_HTTP_ERR"] = NewScCodeEntity(601, 13)
	ScCodeMap["SC_TCP_ERR"] = NewScCodeEntity(602, 13)
	ScCodeMap["SC_UDP_ERR"] = NewScCodeEntity(603, 13)
	ScCodeMap["SC_RPC_ERR"] = NewScCodeEntity(604, 13)
	ScCodeMap["SC_GRPC_ERR"] = NewScCodeEntity(605, 13)
	ScCodeMap["SC_DB_ERR"] = NewScCodeEntity(606, 13)
	ScCodeMap["SC_MQ_ERR"] = NewScCodeEntity(607, 13)
	ScCodeMap["SC_REDIS_ERR"] = NewScCodeEntity(608, 13)
	ScCodeMap["SC_TDENGINE_ERR"] = NewScCodeEntity(609, 13)
	ScCodeMap["SC_FILE_ERR"] = NewScCodeEntity(610, 13)
	ScCodeMap["SC_MYSQL_ERR"] = NewScCodeEntity(611, 13)
	ScCodeMap["SC_NATS_ERR"] = NewScCodeEntity(612, 13)
}

func initGrpcStatusMap() {
	GrpcStatusMap = make(map[uint32]*BaseError)

	GrpcStatusMap[0] = SC_OK
	GrpcStatusMap[2] = SC_UNKNOWN_ERR
	GrpcStatusMap[3] = SC_BAD_REQUEST
	GrpcStatusMap[4] = SC_REQUEST_TIMEOUT
	GrpcStatusMap[5] = SC_NOTFOUND_ERR
	GrpcStatusMap[6] = SC_ALREADY_EXISTS_ERR
	GrpcStatusMap[7] = SC_FORBIDDEN
	GrpcStatusMap[8] = SC_RESOURCE_EXHAUSTED_ERR
	GrpcStatusMap[9] = SC_METHOD_NOT_ALLOWED
	GrpcStatusMap[10] = SC_ABORTED_ERR
	GrpcStatusMap[11] = SC_OUTOFRANGE
	GrpcStatusMap[12] = SC_NOT_IMPLEMENTED
	GrpcStatusMap[13] = SC_SERVER_ERROR
	GrpcStatusMap[14] = SC_SERVICE_UNAVAILABLE
	GrpcStatusMap[15] = SC_DATALOSS_ERR
	GrpcStatusMap[16] = SC_UNAUTHORIZED
}

func IsOk(sysErr *BaseError) bool {
	if sysErr == nil {
		return true
	}
	return ScCodeOkMap.Contains(sysErr.Code)
}

func IsOkCode(code string) bool {
	if code == "" {
		return true
	}
	return ScCodeOkMap.Contains(code)
}

func NewScCodeEntity(httpCode int, grpcCode uint32) *ScCode {
	return &ScCode{
		HttpCode: httpCode,
		GrpcCode: grpcCode,
	}
}

func GetHttpStatus(scCode string) int {
	if v, exist := ScCodeMap[scCode]; exist {
		return v.HttpCode
	}
	return DefaultScHttpCode
}

func GetGrpcStatus(scCode string) uint32 {
	if v, exist := ScCodeMap[scCode]; exist {
		return v.GrpcCode
	}
	return DefaultScGrpcCode
}

func GetErrorByCode(scCode string) *BaseError {
	if v, exist := ScCodeErrMap[scCode]; exist {
		return v
	}
	return nil
}

func GetBaseErrFromGrpcStatus(grpcStatus uint32) *BaseError {
	if val, ok := GrpcStatusMap[grpcStatus]; !ok {
		return SC_GRPC_ERR
	} else {
		return val
	}
}

// AddCodeMap 添加自定义的code和httpCode与grpcCode的对应关系
func AddCodeMap(code string, scCode *ScCode) {
	if code == "" || scCode == nil {
		return
	}
	ScCodeMap[code] = scCode
}
