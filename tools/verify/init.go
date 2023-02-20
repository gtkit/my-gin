package verify

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.superjq.com/go-tools/logger"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gitlab.superjq.com/go-tools/goerr"
)

// 定义一个全局翻译器T
var (
	Trans    ut.Translator
	Validate *validator.Validate
)

// 初始化验证并翻译
func TransValidate() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("validator 初始化失败")
		// return nil, fmt.Errorf("validator 初始化失败")
	}
	// 注册一个获取json tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// 翻译
	err := getTrans("zh", v)
	if err != nil {
		logger.Info("初始化验证翻译 error: ", err)
		panic(err)
	}
	Validate = v
	logger.Info("初始化验证翻译 success")

}

func getTrans(locale string, v *validator.Validate) (err error) {
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, Trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, Trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, Trans)
	}
	// err = v.RegisterTranslation("required_if", Trans, func(ut ut.Translator) error {
	// 	return ut.Add("required_if", "{0}为必填字段!", false) // see universal-translator for details
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("required_if", fe.Field())
	// 	return t
	// })
	return err
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// registerTranslator 为自定义字段添加翻译功能
func RegisterTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func Translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// 翻译自定义校验方法
func SelfRegisterTranslation(v *validator.Validate, method string, info string, myFunc validator.Func) (err error) {

	if err = v.RegisterValidation(method, myFunc); err != nil {
		return
	}

	err = v.RegisterTranslation(
		method,
		Trans,
		RegisterTranslator(method, "{0}"+info),
		Translate,
	)
	return
}

// 完善未有的验证方法的翻译
func AddValidationTranslation(v *validator.Validate, method, info string) error {
	return v.RegisterTranslation(
		method,
		Trans,
		RegisterTranslator(method, "{0}"+info),
		Translate,
	)
}

// 普通验证字段错误信息, 字段名, 验证时的error
func ErrorInfo(field string, err error) error {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		return goerr.Wrap(err, "非validator类型错误")
	}
	for _, v := range RemoveTopStruct(errs.Translate(Trans)) {

		return goerr.Custom(field + " " + v)
	}
	return nil

}
