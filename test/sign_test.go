package test_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"my_gin/internal/pkg/sign"
)

type OrderParams map[string]any

func TestMapsign(t *testing.T) {
	tb := time.Now()
	op := make(OrderParams)
	op["client_id"] = "9aff19ba86e547159d9f1ecc3c322fbb"
	op["access_token"] = "67703ba8c6f143a0834071f6c1372b3c2ff7090a"
	op["data_type"] = "JSON"
	op["order_status"] = 3
	op["page"] = 1
	op["page_size"] = 100
	op["timestamp"] = 1661947835
	op["order_status"] = 3
	op["end_updated_at"] = 1661930091
	op["start_updated_at"] = 1661931891
	op["is_lucky_flag"] = 0
	op["refund_status"] = 1
	op["type"] = "pdd.order.number.list.increment.get"

	appSecret := "c1aa4687a52fb22aea8054eee6dc5123bbb3534d"

	signval := sign.MapSign(op, appSecret)

	t.Log("sign---", signval)
	expectSign := "E652DA7CC3DDD2A415C0C4D03CF449A1"
	t.Log("耗时: ", time.Since(tb))
	assert.Equal(t, expectSign, signval)
}

func BenchmarkMapsign(t *testing.B) {

	t.Run("mapsign", func(b *testing.B) {

	},
	)
	tb := time.Now()
	op := make(OrderParams)
	op["client_id"] = "9aff19ba86e547159d9f1ecc3c322fbb"
	op["access_token"] = "67703ba8c6f143a0834071f6c1372b3c2ff7090a"
	op["data_type"] = "JSON"
	op["order_status"] = 3
	op["page"] = 1
	op["page_size"] = 100
	op["timestamp"] = 1661947835
	op["order_status"] = 3
	op["end_updated_at"] = 1661930091
	op["start_updated_at"] = 1661931891
	op["is_lucky_flag"] = 0
	op["refund_status"] = 1
	op["type"] = "pdd.order.number.list.increment.get"

	appSecret := "c1aa4687a52fb22aea8054eee6dc5123bbb3534d"

	sign := sign.MapSign(op, appSecret)

	// t.Log("sign---", sign)
	expectSign := "E652DA7CC3DDD2A415C0C4D03CF449A1"
	t.Log("耗时: ", time.Since(tb))
	assert.Equal(t, expectSign, sign)
}

func BenchmarkCompare(b *testing.B) {

	b.Run("mapsign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			op := make(OrderParams)
			op["client_id"] = "9aff19ba86e547159d9f1ecc3c322fbb"
			op["access_token"] = "67703ba8c6f143a0834071f6c1372b3c2ff7090a"
			op["data_type"] = "JSON"
			op["order_status"] = 3
			op["page"] = 1
			op["page_size"] = 100
			op["timestamp"] = 1661947835
			op["order_status"] = 3
			op["end_updated_at"] = 1661930091
			op["start_updated_at"] = 1661931891
			op["is_lucky_flag"] = 0
			op["refund_status"] = 1
			op["type"] = "pdd.order.number.list.increment.get"

			appSecret := "c1aa4687a52fb22aea8054eee6dc5123bbb3534d"

			sign.MapSign(op, appSecret)
		}
	})

}

func TestHexiaoMapSgin(t *testing.T) {
	vp := make(OrderParams)
	// now := time.Now()
	vp["type"] = "pdd.voucher.realtime.verify.sync"
	vp["data_type"] = "JSON"
	vp["client_id"] = "862d3b5206b24b018cbf6ca8755e9037"
	vp["access_token"] = "703e5f43b93d4c38b0327cd64e1f172cc4d70394"
	vp["timestamp"] = 1664432742
	vp["request"] = map[string]any{
		"order_sn":       "220825-335502416882023",
		"out_voucher_id": "1123",
		"serial_no":      "12633250409353216",
		"verify_time":    1664432742612,
	}
	appSecret := "2f174ade049c550fce1a813af5a3ada96d8e7fba"

	sign := sign.MapSign(vp, appSecret)
	// expectSign := "763173652C274360B1B5E9FDDEAF2FB0"
	// assert.Equal(t, expectSign, sign)
	t.Log("sign---", sign)
}
