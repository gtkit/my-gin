package test

import (
	"slices"
	"testing"
)

func TestSlices(t *testing.T) {
	s1 := []int{1, 6, 7, 4, 5}
	s2 := slices.Replace(s1, 1, 3, 2)
	t.Log(s1)
	t.Log(s2)

	// for i := range 10 {
	// 	t.Log(i)
	// }
}
