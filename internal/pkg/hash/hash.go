// Package hash 哈希操作类
package hash

import (
	"github.com/gtkit/logger"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptLen  = 60
	bcryptCost = 14
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == bcryptLen
}
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {

		return false, err
	}
	return true, nil
}
