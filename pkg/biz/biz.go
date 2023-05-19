package biz

import (
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// Result Data returns a generic JSON data structure
type Result struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	Redirect string      `json:"redirect,omitempty"`
}

func NewCode(code int, message string, detail ...interface{}) gcode.Code {
	if len(detail) > 0 {
		return gcode.New(code, message, detail[0])
	} else {
		return gcode.New(code, message, nil)
	}

}

func NewError(text ...string) error {
	return gerror.NewCodeSkip(CodeError, 1, text...)
}

func NewErrorf(format string, text ...interface{}) error {
	return gerror.NewCodeSkipf(CodeError, 1, format, text...)
}

func WarpError(err error) error {
	return gerror.NewCodeSkip(CodeError, 1, err.Error())
}

func New() *Result {
	return &Result{}
}

// IsOK Check if the result is CodeOK
func (resp Result) IsOK() bool {
	return resp.Code == CodeOK.Code()
}

// DataString Get Data to string
func (resp Result) DataString() string {
	return gconv.String(resp.Data)
}

// DataInt Get Data as integer
func (resp Result) DataInt() int {
	return gconv.Int(resp.Data)
}

// GetString Get the value of Data as a string by key.
func (resp Result) GetString(key string) string {
	return gconv.String(resp.Get(key))
}

// GetInt Get the value of Data as an integer by key.
func (resp Result) GetInt(key string) int {
	return gconv.Int(resp.Get(key))
}

// Get the value of Data by key, returns a *gvar.Var.
func (resp Result) Get(key string) *gvar.Var {
	m := gconv.Map(resp.Data)
	if m == nil {
		return nil
	}
	return gvar.New(m[key])
}

func (resp Result) JSON() string {
	str, _ := json.Marshal(resp)
	return string(str)
}

// OK return Get general Result
func OK(data ...interface{}) *Result {
	if len(data) > 0 {
		return &Result{Code: CodeOK.Code(), Msg: CodeOK.Message(), Data: data[0]}
	} else {
		return &Result{Code: CodeOK.Code(), Msg: CodeOK.Message()}
	}
}

// Error Get general error Result
func Error(msg ...string) *Result {
	if len(msg) > 0 {
		return &Result{Code: CodeError.Code(), Msg: msg[0]}
	} else {
		return &Result{Code: CodeError.Code(), Msg: CodeError.Message()}
	}
}

// WithCode Get Result with gcode.Code Result
func WithCode(code gcode.Code, msg ...string) *Result {
	message := code.Message()
	if code.Detail() != nil {
		message = code.Detail().(string)
	}
	if len(msg) > 0 {
		message = msg[0]
	}
	return &Result{Code: code.Code(), Msg: message}
}

// Unauthorized Get Authentication failed Result
func Unauthorized(msg ...string) *Result {
	if len(msg) > 0 {
		return &Result{Code: CodeUnauthorized.Code(), Msg: msg[0]}
	} else {
		return &Result{Code: CodeUnauthorized.Code(), Msg: CodeUnauthorized.Message()}
	}
}

// Forbidden Get No Access Result
func Forbidden(msg ...string) *Result {
	if len(msg) > 0 {
		return &Result{Code: CodeForbidden.Code(), Msg: msg[0]}
	} else {
		return &Result{Code: CodeForbidden.Code(), Msg: CodeForbidden.Message()}
	}
}

// BadRequest Get Authentication failed Result
func BadRequest(msg ...string) *Result {
	if len(msg) > 0 {
		return &Result{Code: CodeBadRequest.Code(), Msg: msg[0]}
	} else {
		return &Result{Code: CodeBadRequest.Code(), Msg: CodeBadRequest.Message()}
	}
}
