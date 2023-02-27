package resty

import (
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

var Rclient *resty.Client

func NewClient() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	client := resty.New()
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal
	Rclient = client
}
