package validator

import "github.com/go-playground/validator/v10"

type (
	//参考：https://echo.labstack.com/guide/request
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// todo 验证
func New() (v *CustomValidator) {
	v = &CustomValidator{validator: validator.New()}
	return
}
