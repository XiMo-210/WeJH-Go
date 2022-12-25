package apiException

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	OpenIDError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError            = NewError(http.StatusInternalServerError, 200501, "参数错误")
	NotAdmin              = NewError(http.StatusInternalServerError, 200502, "该用户无管理员权限")
	UserNotFind           = NewError(http.StatusInternalServerError, 200503, "该用户不存在")
	NotLogin              = NewError(http.StatusInternalServerError, 200503, "未登录")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200504, "密码错误")
	HttpTimeout           = NewError(http.StatusInternalServerError, 200505, "系统异常，请稍后重试!")
	RequestError          = NewError(http.StatusInternalServerError, 200506, "系统异常，请稍后重试!")
	NotBindYxy            = NewError(http.StatusInternalServerError, 200507, "该用户未绑定或未注册易校园")
	UserAlreadyExisted    = NewError(http.StatusInternalServerError, 200508, "该用户已激活")
	WrongVerificationCode = NewError(http.StatusInternalServerError, 200509, "验证码错误")
	StudentNumAndIidError = NewError(http.StatusInternalServerError, 200510, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError              = NewError(http.StatusInternalServerError, 200511, "密码长度必须在6~20位之间")
	ReactiveError         = NewError(http.StatusInternalServerError, 200512, "该通行证已经存在，请重新输入")
	StudentIdError        = NewError(http.StatusInternalServerError, 200513, "学号格式不正确，请重新输入")
	NotInit               = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	NotFound              = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	Unknown               = NewError(http.StatusInternalServerError, 300500, "系统异常，请稍后重试!")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}
