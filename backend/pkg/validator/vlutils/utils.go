package vlutils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ErrTranslationsToStr(errs validator.ValidationErrorsTranslations) string {
	errList := make([]string, 0, len(errs))
	for _, err := range errs {
		errList = append(errList, err)
	}

	return strings.Join(errList, ", ")
}
