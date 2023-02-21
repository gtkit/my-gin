package response

import (
	"fmt"
	"net/http"
	"sync"
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
func NewResponse(status int, code uint32, data interface{}) *Response {
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.Res.Code = code
	response.Res.Message = Text(code)
	response.Res.Data = data

	return response
}

func PutResponse(r *Response) {
	pool.Put(r)
}

func ResponseOk(code uint32, data interface{}) *Response {
	return NewResponse(http.StatusOK, code, data)
}
