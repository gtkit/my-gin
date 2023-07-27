package response

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

type resResult struct {
	Code    ErrCode     `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

// 为了提高效率使用一个Pool
var pool = sync.Pool{
	New: func() interface{} {
		return &resResult{}
	},
}

// 定义自己的返回code
func NewResponse(code ErrCode, msg string, data, meta interface{}) *resResult {
	response := pool.Get().(*resResult)

	response.Code = code
	response.Message = msg
	response.Data = data
	response.Meta = meta

	return response
}

func Ok(c *gin.Context, data interface{}) {
	res := NewResponse(Success, Success.String(), data, nil)
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func OkWithMeta(c *gin.Context, data, meta interface{}) {
	res := NewResponse(Success, Success.String(), data, meta)
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Err(c *gin.Context, err goerr.Error) {
	res := NewResponse(ErrCode(err.Code().Code), err.Error(), [0]int{}, nil)
	c.SecureJSON(err.Code().HTTPCode, res)
	PutResponse(res)
}

func Fail(c *gin.Context) {
	res := NewResponse(RequestFail, RequestFail.String(), [0]int{}, nil)
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Error(c *gin.Context, err goerr.Error) {
	res := NewResponse(ErrCode(err.Code().Code), err.Error(), [0]int{}, nil)
	c.SecureJSON(err.Code().HTTPCode, res)
	PutResponse(res)
}

func NotFoundError(c *gin.Context, errmsg string) {
	res := NewResponse(http.StatusNotFound, errmsg, [0]int{}, nil)
	c.SecureJSON(http.StatusNotFound, res)
	PutResponse(res)
}

func PutResponse(r *resResult) {
	r.Code = 0
	r.Message = ""
	r.Data = nil
	r.Meta = nil

	pool.Put(r)
}
