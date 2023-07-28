package resty

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gtkit/logger"
	jsoniter "github.com/json-iterator/go"
)

var restyClient *resty.Client

func NewClient() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	client := resty.New()
	client.SetTimeout(5 * time.Second)                               // 设置超时时间为 5 秒钟
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // 关闭证书校验
	client.SetRetryCount(3).SetRetryWaitTime(1 * time.Second)        // 设置最大重试次数为 3 次，重试间隔时间为 1 秒钟
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal
	client.SetLogger(&restyLogger{})

	restyClient = client
}

func Client() *resty.Client {
	return restyClient
}

func Request() *resty.Request {
	return restyClient.R()
}

// restyLogger resty logger
var _ resty.Logger = (*restyLogger)(nil)

type restyLogger struct {
}

func (l *restyLogger) Errorf(format string, v ...interface{}) {
	logger.Errorf("--ERROR RESTY "+format, v)
}
func (l *restyLogger) Warnf(format string, v ...interface{}) {
	logger.Warnf("--WARN RESTY "+format, v)
}
func (l *restyLogger) Debugf(format string, v ...interface{}) {
	logger.Debugf("--DEBUG RESTY "+format, v)
}

// 使用教程: https://blog.csdn.net/qq_29799655/article/details/130831278
