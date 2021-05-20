package validators

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
	Trans    ut.Translator
)

func InitValidator() {

	zhs := zh.New()
	Uni = ut.New(zhs)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	Trans, _ = Uni.GetTranslator("zh")

	Validate = validator.New()
	zhTranslations.RegisterDefaultTranslations(Validate, Trans)

}

// 返回错误信息
func GetErrorMessage(err error) *string {

	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	var errMessage string
	for _, e := range errs {
		// can translate each error one at a time.
		errMessage = errMessage + e.Translate(Trans)

	}
	return &errMessage
}
