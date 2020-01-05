package ykit

/*--auth: whr  date:2019/12/511:43--------------------------
 ####请勿擅改此功能代码####
 用途：
 统一的服务返回数据接口，用于标准化输出
--------------------------------------- */
type (
	Result struct {
		//正常返回值为200，
		//小于0为错误，大于0的非200值为有警告
		Code int `json:"code"`
		//文本消息，错误内容或警告内容，可执行成功
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

const (
	OK      = 200
	FAIL    = -1
	WARNING = 1000
)

func RErr(msg string) *Result {
	return &Result{
		Code: FAIL,
		Msg:  msg,
		Data: nil,
	}
}

func ROK(msg string, data interface{}) *Result {
	return &Result{
		Code: OK,
		Msg:  msg,
		Data: data,
	}
}

func RWarn(msg string, data interface{}) *Result {
	return &Result{
		Code: WARNING,
		Msg:  msg,
		Data: data,
	}
}

func ResultOK(data interface{}) *Result {
	return &Result{
		Code: 200,
		Msg:  "执行成功",
		Data: data,
	}
}

func ResultError(err error) *Result {
	return &Result{
		Code: -1,
		Msg:  err.Error(),
		Data: nil,
	}
}

func ResultWarning(warnMsg string, data interface{}) *Result {
	return &Result{
		Code: 100,
		Msg:  warnMsg,
		Data: data,
	}
}
