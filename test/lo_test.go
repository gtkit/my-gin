// @Author xiaozhaofu 2023/3/1 20:07:00
package test_test

import (
	"testing"

	"github.com/samber/lo"

	"ydsd_gin/tools/utils"
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

func TestSliceLo(t *testing.T) {
	a := make([]int, 0)
	a = append(a, 1)
	ie := lo.IsEmpty[int](len(a))
	t.Log("ie slice---", ie)
}

func TestInterfaceLo(t *testing.T) {
	var a interface{}
	// a = make([]int, 0)
	// a = make(map[string]string, 0)
	t.Log("a is nil--", a == nil)
	ie := lo.IsEmpty(a)
	t.Log("ie interface---", ie)

	ia := utils.Empty(a)
	t.Log("ia interface----", ia)
}
