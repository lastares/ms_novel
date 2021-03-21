package boot

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"ms_novel/global"
	"os"
	"reflect"
)

func validateInit() {
	// 实例化全局验证器
	global.Validate = validator.New()
	// 实例化中文翻译插件实例
	universal := ut.New(zh.New())
	// 获取翻译插件类型
	translator, _ := universal.GetTranslator("zh")
	// 给全局验证器注册一个翻译插件
	err := zh_translation.RegisterDefaultTranslations(global.Validate, translator)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	// 注册函数，获取 struct 结构体中的中文字段名，这里定义label作为字段名，当然也可以是其他的，顺眼就好
	global.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 全局翻译插件
	global.Translator = translator
}
