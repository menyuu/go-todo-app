package forms

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// バリデーションの作成
func ValidateStruct(form interface{}) map[string]string {
	errors := make(map[string]string)

	err := validate.Struct(form)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = validationMessage(e)
		}
	}
	return errors
}

// バリデーションメッセージの作成
func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + "は必須です"
	case "email":
		return "正しいメールアドレスを入力してください"
	case "min":
		return e.Field() + "は最小" + e.Param() + "文字必要です"
	case "max":
		return e.Field() + "は最大" + e.Param() + "文字必要です"
	default:
		return "入力内容に誤りがあります"
	}
}
