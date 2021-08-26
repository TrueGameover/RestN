package normalizer

import "github.com/TrueGameover/RestN/rest"

type FieldValidationErrorNormalizer struct {
	rest.IResponseNormalizer
}

func (n FieldValidationErrorNormalizer) Normalize(object interface{}, _ rest.NormalizeMethod, _ rest.Options, _ int) interface{} {
	fe, _ := object.(rest.FieldValidationError)

	dict := map[string]interface{}{
		"Field":   fe.Field,
		"Message": fe.Message,
	}

	return &dict
}

func (n FieldValidationErrorNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.FieldValidationError)
	return
}
