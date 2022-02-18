package errcode

import (
	"fmt"
	"net/http"
)

// 错误处理：错误处理公共方法，用以标准化项目的错误输出

type Error struct {
	// 1.首先声明Error结构体，用于表示错误的响应结果
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"detail"`
}

var codes = map[int]string{}

// 2.把codes作为全局错误码的存储载体，以便查看当前的注册情况

func NewError(code int, msg string) *Error {
	// 3.在调用NewError创建新的Error实例的同时，进行排重校验
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	// 4.比较特殊的是StatusCode方法，它主要针对一些特定错误码进行状态码的转换。因为不同的内部错误码在HTTP状态码中表示不同的含义，
	// 所以我们需要将其区分开来，以便客户端及监控或报警等系统的识别和监听。
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthhorizedAuthNotExist.Code():
		fallthrough
	case UnauthhorizedTokenError.Code():
		fallthrough
	case UnauthhorizedTokenGenerate.Code():
		fallthrough
	case UnauthhorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
