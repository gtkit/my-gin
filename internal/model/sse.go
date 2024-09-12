// @Author xiaozhaofu 2023/5/29 21:12:00
package model

import (
	"context"
)

//go:generate go-option -type SseClient
type SseClient struct {
	Ctx     context.Context // Connection的上下文
	Message chan string
	Cancel  context.CancelFunc // 取消Connection的方法
	Err     chan error
	// This is the channel used to send messages to client
	Id string
}

type SseMsg struct {
	Msg  string `json:"msg"`
	Code int32  `json:"code"`
}
