// @Author xiaozhaofu 2023/3/1 20:07:00
package test

import (
	"testing"

	"github.com/samber/lo"
)

func TestStrLo(t *testing.T) {
	str := ""
	ie := lo.IsEmpty(str)
	t.Log("ie str----", ie)
}

func TestNumLo(t *testing.T) {
	num := 0
	ie := lo.IsEmpty(num)
	t.Log("ie num----", ie)
}
