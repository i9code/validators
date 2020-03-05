package validators

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type UserInfo struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required,check_cn_sp"`
}

func TestNewValidate(t *testing.T) {
	validate := NewValidate()
	user := &UserInfo{
		FirstName: "我是中国人，我爱自己的祖国",
		LastName:  "A aacd",
	}
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}
}
