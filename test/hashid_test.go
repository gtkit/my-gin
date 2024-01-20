// @Author xiaozhaofu 2023/2/28 19:32:00
package test_test

import (
	"testing"

	"github.com/gtkit/encry/hids"
	"github.com/magiconair/properties/assert"
)

func TestHashid(t *testing.T) {
	sec := "6ab6122836Tfef95f8$b"
	ids := []int{18}
	l := 12
	h := hids.New(sec, l)
	enid, err := h.EncodeHashids(ids)
	if err != nil {
		t.Error("enid err----", err)
	}
	t.Log("enid-----", enid)
	deid, err := h.DecodeHashids(enid)
	if err != nil {
		t.Error("deid err-----", err)
	}
	t.Log("deid-----", deid)
	assert.Equal(t, deid, ids)
}
