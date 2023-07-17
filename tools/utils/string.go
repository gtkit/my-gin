package utils

import (
	"crypto/rand"
	"io"
	"log"
	mathrand "math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// IsBlank checks whether the given string is blank.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {

		return false, err
	}
	return true, nil
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}
func StrToInt(val string) int {
	v1, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return int(v1)
}
func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func B2I(bs []uint8) (int, error) {
	i, err := strconv.Atoi(B2S(bs))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func B2I64(bs []uint8) (int64, error) {
	i, err := strconv.ParseInt(B2S(bs), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
