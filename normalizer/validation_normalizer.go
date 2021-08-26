package normalizer

import "github.com/TrueGameover/RestN/rest"

type ValidationNormalizer struct {
	rest.IResponseNormalizer
}

func (v ValidationNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	val, _ := object.(rest.Validation)

	dict := map[string]interface{}{
		"FieldErrors": normalize(val.FieldErrors, options, depth),
	}

	return &dict
}

func (v ValidationNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.Validation)
	return
}
