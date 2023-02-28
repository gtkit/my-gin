// @Author xiaozhaofu 2023/2/28 19:32:00
package test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"ydsd_gin/tools/hash"
)

func TestHashid(t *testing.T) {
	sec := "123456"
	ids := []int{18}
	l := 12
	h := hash.New(sec, l)
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
