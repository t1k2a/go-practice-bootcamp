package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		// 💡 チャレンジ課題（任意）
		// エラー文言を "title is required" のように見やすく整形
		errors[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
	}

	return errors
}