package response

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"
)

//go:generate stringer -type ErrCode -linecomment

type ErrCode int // 错误码

const (
	Success     ErrCode = 200 // success
	RequestFail ErrCode = 101 // 请求失败
)

type resResult struct {
	Code    ErrCode     `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 为了提高效率使用一个Pool
var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("----new Response pool----")
		return &resResult{}
	},
}

// 定义自己的返回code
func NewResponse(code ErrCode, msg string, data interface{}) *resResult {
	response := pool.Get().(*resResult)

	response.Code = code
	response.Message = msg
	response.Data = data

	return response
}

func Ok(c *gin.Context, data interface{}) {
	res := NewResponse(Success, Success.String(), data)
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Fail(c *gin.Context) {
	res := NewResponse(RequestFail, RequestFail.String(), [0]int{})
	c.SecureJSON(http.StatusOK, res)
	PutResponse(res)
}

func Error(c *gin.Context, err goerr.Error) {
	res := NewResponse(ErrCode(err.Code().Code), err.Error(), [0]int{})
	c.SecureJSON(err.Code().HTTPCode, res)
	PutResponse(res)
}

func NotFoundError(c *gin.Context, errmsg string) {
	res := NewResponse(http.StatusNotFound, errmsg, [0]int{})
	c.SecureJSON(http.StatusNotFound, res)
	PutResponse(res)
}

func PutResponse(r *resResult) {
	// fmt.Println("------put response pool---")
	r.Code = 0
	r.Message = ""
	r.Data = nil

	pool.Put(r)
}
