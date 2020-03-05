package validators

import (
	"regexp"
	"github.com/go-playground/validator/v10"
	"strings"
)

func NewValidate()*validator.Validate{
	validate:= validator.New()
	//注册自定义函数，与struct tag关联起来
	err := validate.RegisterValidation("check_cnsp", CheckStringIsCnOrUpperAndNoSpecialCharacter)
	if err != nil {
		panic("注册 英文单词首字母必须大写 只能含有中文和英文 不能包含标点符号如_ 检测失败")
	}
	return validate
}

var mstr string = `^([A-Z]|[\x{2e80}-\x{2e99}\x{2e9b}-\x{2ef3}\x{2f00}-\x{2fd5}\x{3005}-\x{3007}\x{3021}-\x{3029}\x{3038}-\x{303b}\x{3400}-\x{4db5}\x{4e00}-\x{9fef}\x{f900}-\x{fa6d}\x{fa70}-\x{fad9}\x{20000}-\x{2a6d6}\x{2a700}-\x{2b734}\x{2b740}-\x{2b81d}\x{2b820}-\x{2cea1}\x{2ceb0}-\x{2ebe0}\x{2f800}-\x{2fa1d}])[a-zA-Z0-9\x{2e80}-\x{2e99}\x{2e9b}-\x{2ef3}\x{2f00}-\x{2fd5}\x{3005}-\x{3007}\x{3021}-\x{3029}\x{3038}-\x{303b}\x{3400}-\x{4db5}\x{4e00}-\x{9fef}\x{f900}-\x{fa6d}\x{fa70}-\x{fad9}\x{20000}-\x{2a6d6}\x{2a700}-\x{2b734}\x{2b740}-\x{2b81d}\x{2b820}-\x{2cea1}\x{2ceb0}-\x{2ebe0}\x{2f800}-\x{2fa1d}]*$`

/* 所有英文单词首字母必须大写 只能含有中文和英文 不能包含标点符号如_,-等 尽可能支持国际化*/
func CheckStringIsCnOrUpperAndNoSpecialCharacter(fl validator.FieldLevel)  (ret bool) {
	ret = true
	strs := strings.Fields(fl.Field().String())
	for _, s := range strs {
		if ok, _ := regexp.MatchString(mstr, s); !ok {
			ret = ok
			return
		}
	}
	return
}