// @Author xiaozhaofu 2023/3/1 14:56:00
package test_test

import (
	"testing"

	"my_gin/tools/utils"
)

func TestString(t *testing.T) {
	t.Log(utils.RandomString(6))
}

func TestNumber(t *testing.T) {
	t.Log(utils.RandomNumber(12))
}
