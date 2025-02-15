package config

import (
	"Democratic_shopping_mall/global"

	"github.com/go-playground/validator/v10"
)

func InitValidate() {
	global.Validate = validator.New()
}
