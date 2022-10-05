package validation

import (
	"github.com/essaherlandy/absensi/app/models"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
)

func ValidateStruct(user models.User) map[string]string {
	var errors = make(map[string]string)

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	var validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, err := range errs {
			errors[err.Field()] = err.Translate(trans)
		}
	}

	return errors
}
