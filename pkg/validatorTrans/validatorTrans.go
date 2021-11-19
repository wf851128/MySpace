package validatorTrans

import (
	"MySpace/models"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// Trans 定义一个全局的翻译器
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locals string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取 json tag 引擎属性，实现自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		//注册 SignUpParam 注册自定义校验方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation, models.ParamSignUp{})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用（fallback）的语言环境
		//第二个和第三个参数是应该支持的语言环境（支持多个）
		uni := ut.New(enT, zhT, enT)
		//locale 取决于 http 请求头的`Accept-Language`
		var ok bool
		//也可以使用 uni.FindTranslator(...)传入多个 locale 来进行查找
		Trans, ok = uni.GetTranslator(locals)
		if !ok {
			return fmt.Errorf("controllers.SignUpHandler.ShouldBindQuery(%s) failed", locals)
		}
		//	注册翻译器
		switch locals {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		if err != nil {
			return fmt.Errorf("controllers.SignUpHandler.RegisterDefaultTranslations(%s) failed", locals)
		}
		return
	}

	return
}

// RemoveTopStruct 去除提示信息中的结构体名称，可能影响效率
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// SignUpParamStructLevelValidation  自定义校验函数，对注册信息进行自定义校验
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSignUp)
	if su.Password != su.RePassword {
		sl.ReportError(su.RePassword,
			"re_password",
			"RePassword",
			"eqfield",
			"password")
	}
}
