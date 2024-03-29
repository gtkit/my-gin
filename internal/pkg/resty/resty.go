package resty

import (
	"crypto/tls"
	"time"

	"github.com/gtkit/logger"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	timeout       = 5 * time.Second
	tryNum        = 3
	retryInterval = 1
)

var restyClient *resty.Client

func NewClient() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	client := resty.New()
	client.SetTimeout(timeout)                                       // 设置超时时间为 5 秒钟
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // 关闭证书校验
	client.SetRetryCount(tryNum).SetRetryWaitTime(retryInterval)     // 设置最大重试次数为 3 次，重试间隔时间为 1 秒钟
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal
	client.SetLogger(logger.RestyLogger())

	restyClient = client
}

func Client() *resty.Client {
	return restyClient
}

func R() *resty.Request {
	return restyClient.R()
}

// 使用教程: https://blog.csdn.net/qq_29799655/article/details/130831278
