// @Author xiaozhaofu 2023/5/29 21:12:00
package model

import (
	"context"
)

type SseClient struct {
	// This is the channel used to send messages to client
	Id      string
	Message chan string
	Ctx     context.Context    // Connection的上下文
	Cancel  context.CancelFunc // 取消Connection的方法
	Err     chan error
}

type SseMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
