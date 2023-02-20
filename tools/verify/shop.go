package verify

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"ydsd_gin/internal/model"
)

// customFunc 自定义字段级别校验方法,验证日期要在当前日期后
func CustomFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	fmt.Println("获取到的日期")
	fmt.Println(date)
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// 自定义验证函数
func CheckName(fl validator.FieldLevel) bool {
	if fl.Field().String() != "roottom" {
		return false
	}
	return true
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(model.SignUpParam)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "填入的password")
	}
}
