package resp

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"
)

//go:generate stringer -type ErrCode -linecomment

type ErrCode int // 错误码

const (
	Success     ErrCode = 200      // success
	RequestFail ErrCode = 10010014 // 请求失败
)

type ResResult struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"msg"`
	Data    any     `json:"data"`
	Meta    any     `json:"meta,omitempty"`
}

// 为了提高效率使用一个Pool
var pool = sync.Pool{
	New: func() any {
		return &ResResult{}
	},
}

// 定义自己的返回code
func NewResponse(code ErrCode, msg string, data, meta any) *ResResult {
	response, ok := pool.Get().(*ResResult)
	if !ok {
		return nil
	}

	response.Code = code
	response.Message = msg
	response.Data = data
	response.Meta = meta

	return response
}

func Ok(c *gin.Context, data any) {
	res := NewResponse(Success, Success.String(), data, nil)
	c.Abort()
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func OkWithMeta(c *gin.Context, data, meta any) {
	res := NewResponse(Success, Success.String(), data, meta)
	c.Abort()
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Err(c *gin.Context, err goerr.Error) {
	res := NewResponse(ErrCode(err.ErrCode()), err.Error(), [0]int{}, nil)
	c.SecureJSON(err.HttpCode(), res)
	PutResponse(res)
}

func Fail(c *gin.Context) {
	res := NewResponse(RequestFail, RequestFail.String(), [0]int{}, nil)
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Error(c *gin.Context, err goerr.Error) {
	res := NewResponse(ErrCode(err.ErrCode()), err.Error(), [0]int{}, nil)
	c.SecureJSON(err.HttpCode(), res)
	PutResponse(res)
}

func NotFoundError(c *gin.Context, errmsg string) {
	res := NewResponse(http.StatusNotFound, errmsg, [0]int{}, nil)
	c.SecureJSON(http.StatusNotFound, res)
	PutResponse(res)
}

func NotAllowedMethod(c *gin.Context, errmsg string) {
	res := NewResponse(http.StatusMethodNotAllowed, errmsg, [0]int{}, nil)
	c.SecureJSON(http.StatusMethodNotAllowed, res)
	PutResponse(res)
}

func PutResponse(r *ResResult) {
	r.Code = 0
	r.Message = ""
	r.Data = nil
	r.Meta = nil

	pool.Put(r)
}
