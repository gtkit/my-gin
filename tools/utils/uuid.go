// @Author xiaozhaofu 2023/3/1 20:25:00
package utils

import "github.com/google/uuid"

func NewUuid() string {
	return uuid.New().String()
}
func UUID() string {
	return uuid.New().String()
}
