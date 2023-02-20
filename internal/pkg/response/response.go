package response

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type resResult struct {
	Code    uint32      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type Response struct {
	HttpStatus int
	Res        resResult
}

// 为了提高效率我们可以使用一个Pool
var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("----new Response pool----")
		return &Response{}
	},
}

// 定义自己的返回code
const (
	QuerySuccess uint32 = iota
	QueryFailed
)

var MessageForCode = map[uint32]string{
	QuerySuccess: "查询成功",
	QueryFailed:  "查询失败",
}

func NewResponse(status int, code uint32, data interface{}) *Response {
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.Res.Code = code
	response.Res.Message = Msg[code]
	response.Res.Data = data

	return response
}

func PutResponse(r *Response) {
	r = &Response{}
	pool.Put(r)
}

func ResponseOk(code uint32, data interface{}) *Response {
	return NewResponse(http.StatusOK, code, data)
}

type Handler func(ctx *gin.Context) *Response

func Decorate(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := h(c)
		if r != nil {
			c.SecureJSON(r.HttpStatus, &r.Res)
		}
		PutResponse(r)
	}
}
