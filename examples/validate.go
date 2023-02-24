package examples

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.superjq.com/go-tools/goerr"

	"ydsd_gin/internal/pkg/response"
	"ydsd_gin/tools/verify"
)

type SignUpParam struct {
	Age        uint8  `json:"age" form:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" form:"name" binding:"contains=tom,checkName"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password" form:"password" binding:"required"`
	RePassword string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
	Date       string `json:"date" form:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

// 校验方式使用 demo
func Shop(c *gin.Context) {
	_ = verify.Validate.Struct(SignUpParam{}) // 执行验证
	verify.Validate.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})

	if err := verify.SelfRegisterTranslation(verify.Validate, "checkDate", "必须要晚于当前日期", CustomFunc); err != nil {
		panic(err)
	}

	if err := verify.SelfRegisterTranslation(verify.Validate, "checkName", "名字格式不对", CheckName); err != nil {
		panic(err)
	}

	var s SignUpParam

	//
	if err := c.ShouldBindQuery(&s); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			// app.Error(c, goerr.ErrParams, err)
			return
		}

		for _, v := range verify.RemoveTopStruct(errs.Translate(verify.Trans)) {
			response.Error(c, goerr.ErrValidateParams, goerr.Custom(v))
			return
		}

	}

	// app.OK(c, "success")

}

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
	su := sl.Current().Interface().(SignUpParam)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "填入的password")
	}
}
func FilterCategory(c *gin.Context) {
	typeid := c.DefaultQuery("type_id", "1")

	tid, err := strconv.Atoi(typeid)
	if err != nil {

		return
	}
	if err = verify.ErrorInfo("type_id", verify.Validate.Var(tid, "ne=2,ne=3,gte=1,lte=11")); err != nil {

		return
	}

}
