// @Author xiaozhaofu 2023/3/17 20:37:00
package test_test

import (
	"encoding/base64"
	"testing"

	"github.com/gtkit/encry/rsa"
	"github.com/magiconair/properties/assert"
)

// 测试解密
func TestDecodeRsa(t *testing.T) {
	want := origin()
	secstr := "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"
	data, err := base64.StdEncoding.DecodeString(secstr)
	if err != nil {
		t.Error("base64.StdEncoding.DecodeString--", err)
	}
	p, err := rsa.RsaDecryptBlock(data, "/Users/xiaozhaofu/go/src/office-lenovo-pay/config/pem/private.pem")
	if err != nil {
		t.Error("rsa.RsaDecrypt err---", err)
	}

	t.Log("origin str--", string(p))
	assert.Equal(t, string(p), want)
}

// 测试加密
func TestEncodeRsa(t *testing.T) {
	origin := origin()
	secstr, err := rsa.RsaEncryptBlock([]byte(origin), "/Users/xiaozhaofu/go/src/office-lenovo-pay/config/pem/public.pem")
	if err != nil {
		t.Error("rsa.RsaEncryptBlock err---", err)
	}
	t.Log("secstr---", secstr)
}

func origin() string {
	return "https://go.microsoft.com/fwlink/?LinkID=529180&aid=f1d4cfd0-c23b-4f35-a9c8-c2bac1ce61be&invite=6bb1b92b-d621-47b6-aade-cdb7471d373aqq&type=ByLink"
}
