package resty

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gtkit/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
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

	// 设置 JSON 序列化器和反序列化器
	client.SetJSONMarshaler(json.Marshal)
	client.SetJSONUnmarshaler(json.Unmarshal)
	client.EnableTrace() // 开启请求跟踪
	// 设置日志
	client.SetLogger(logger.RestyLogger())
	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
			// Do something with v.Response
			logger.ZError("resty error", zap.Error(v.Err))
		}
		// Log the error, increment a metric, etc...
	})

	restyClient = client
}

func Client() *resty.Client {
	return restyClient
}

func R() *resty.Request {
	return restyClient.R()
}

// 使用教程: https://blog.csdn.net/qq_29799655/article/details/130831278

// ProtobufMarshaller 定义Protobuf的Marshaller
func ProtobufMarshaller(content any) ([]byte, error) {
	msg, ok := content.(proto.Message)
	if !ok {
		return nil, errors.New("content is not a proto.Message")
	}
	return proto.Marshal(msg)
}

// ProtobufUnmarshaller 定义Protobuf的Unmarshaller
func ProtobufUnmarshaller(data []byte, content any) error {
	return proto.Unmarshal(data, content.(proto.Message))
}
