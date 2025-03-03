package validator

import (
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	rutranslations "github.com/go-playground/validator/v10/translations/ru"
)

type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func NewRuValidator() *Validator {
	validate := validator.New()

	russian := ru.New()
	uni := ut.New(russian, russian)
	trans, _ := uni.GetTranslator("ru")

	_ = rutranslations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		Validate: validate,
		Trans:    trans,
	}
}
