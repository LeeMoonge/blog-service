package errcode

// 公共错误码：预定义项目中的一些公共错误码，以便引导和规范使用

var (
	Success                    = NewError(0, "成功")
	ServerError                = NewError(10000000, "服务内部错误")
	InvalidParams              = NewError(10000001, "入参错误")
	NotFound                   = NewError(10000002, "找不到")
	UnauthhorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthhorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthhorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthhorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests            = NewError(10000007, "请求过多")
)
