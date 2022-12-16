package app

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

func NewCustomValidator() (v *CustomValidator) {
	v = &CustomValidator{validator: validator.New()}
	return
}
