package errcode

var (
	Success                    = NewError(0, "success")
	ServerError                = NewError(-1, "server error")
	InvalidParams              = NewError(-2, "invalid params")
	NotFound                   = NewError(-3, "not found")
	UnauthorizedAuthNotExist   = NewError(-4, "unauthorized, auth not exist")
	UnauthorizedTokenError     = NewError(-5, "unauthorized, token error")
	UnauthorizedTokenTimeout   = NewError(-6, "unauthorized, token timeout")
	UnauthorizedTokenGenerate  = NewError(-7, "unauthorized, token generate error")
	TooManyRequests            = NewError(-8, "too many requests")
	UnauthorizedAuthTokenError = NewError(-9, "unauthorized, auth token error")
)
