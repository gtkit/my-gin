package test_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSize(t *testing.T) {
	fmt.Println(unsafe.Sizeof(int(0)))   // 8
	fmt.Println(unsafe.Sizeof(int8(0)))  // 1
	fmt.Println(unsafe.Sizeof(int16(0))) // 2
	fmt.Println(unsafe.Sizeof(int32(0))) // 4
	fmt.Println(unsafe.Sizeof(int64(0))) // 8

	fmt.Println(unsafe.Sizeof(uint(0)))   // 8
	fmt.Println(unsafe.Sizeof(uint8(0)))  // 1
	fmt.Println(unsafe.Sizeof(uint16(0))) // 2
	fmt.Println(unsafe.Sizeof(uint32(0))) // 4
	fmt.Println(unsafe.Sizeof(uint64(0))) // 8

	fmt.Println(unsafe.Sizeof(byte(0)))       // 1
	fmt.Println(unsafe.Sizeof(rune(0)))       // 4
	fmt.Println(unsafe.Sizeof(uintptr(0)))    // 8
	fmt.Println(unsafe.Sizeof(float32(0)))    // 4
	fmt.Println(unsafe.Sizeof(float64(0)))    // 8
	fmt.Println(unsafe.Sizeof(complex64(0)))  // 8
	fmt.Println(unsafe.Sizeof(complex128(0))) // 16

	fmt.Println(unsafe.Sizeof(false))    // 1
	fmt.Println(unsafe.Sizeof("string")) // 16
}

type Student struct {
	height uint8
	age    uint8
}

func TestSize2(t *testing.T) {
	arr := [5]uint8{}               // array
	fmt.Println(unsafe.Sizeof(arr)) // 5

	stu := Student{height: 175, age: 36} // struct
	fmt.Println(unsafe.Sizeof(stu))      // 2

	stuA := new(Student)             // pointer
	fmt.Println(unsafe.Sizeof(stuA)) // 8

	stuB := []Student{stu}                      // slice
	fmt.Println(unsafe.Sizeof(stuB))            // 24
	stuB2 := []Student{stu, stu, stu, stu, stu} // slice
	fmt.Println(unsafe.Sizeof(stuB2))           // 24

	stuC := make(map[string]Student)      // map
	fmt.Println(unsafe.Sizeof(stuC))      // 8
	stuC2 := make(map[string]Student, 64) // map
	fmt.Println(unsafe.Sizeof(stuC2))     // 8

	stuD := make(chan Student)        // channel
	fmt.Println(unsafe.Sizeof(stuD))  // 8
	stuD2 := make(chan Student, 32)   // channel
	fmt.Println(unsafe.Sizeof(stuD2)) // 8

	stuE := make([]Student, 8)        // slice
	fmt.Println(unsafe.Sizeof(stuE))  // 24
	stuE2 := make([]Student, 16, 32)  // slice
	fmt.Println(unsafe.Sizeof(stuE2)) // 24

	stuF0 := new([]Student)           // pointer of a empty slice
	fmt.Println(unsafe.Sizeof(stuF0)) // 8

	var emptyFunc = func() {}             // empty func
	fmt.Println(unsafe.Sizeof(emptyFunc)) // 8

	var stuF1 []Student               // empty slice
	fmt.Println(unsafe.Sizeof(stuF1)) // 24
	var stuF2 struct{}                // empty struct
	fmt.Println(unsafe.Sizeof(stuF2)) // 0
	var stuF3 [0]Student              // empty array
	fmt.Println(unsafe.Sizeof(stuF3)) // 0
	var stuF4 any                     // empty interface
	fmt.Println(unsafe.Sizeof(stuF4)) // 16
}
