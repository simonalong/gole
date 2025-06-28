package errorx

import "fmt"

type BaseError struct {
	error
	Code   string `json:"code"`
	Msg    string `json:"msg,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func New(code, message string) *BaseError {
	return &BaseError{
		Code: code,
		Msg:  message,
	}
}

func NewWithAll(code, message, detail string) *BaseError {
	return &BaseError{
		Code:   code,
		Msg:    message,
		Detail: detail,
	}
}

func NewWithDetail(code, detail string) *BaseError {
	baseErr := GetErrorByCode(code)
	if baseErr == nil {
		return &BaseError{
			Code:   code,
			Msg:    "未知异常",
			Detail: detail,
		}
	} else {
		baseErr.Detail = detail
		return baseErr
	}
}

func (e *BaseError) Error() string {
	if e.Detail == "" {
		return fmt.Sprintf("code=%s, msg=%s", e.Code, e.Msg)
	}
	return fmt.Sprintf("code=%s, msg=%s, detail=【%s】", e.Code, e.Msg, e.Detail)
}

func (e *BaseError) WithMsg(msg string) *BaseError {
	e.Msg = msg
	return e
}

func (e *BaseError) WithError(err error) *BaseError {
	e.Detail = err.Error()
	return e
}

func (e *BaseError) WithDetail(detail string) *BaseError {
	e.Detail = detail
	return e
}
