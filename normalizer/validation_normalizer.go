package normalizer

import "RestN/rest"

type ValidationNormalizer struct {
	rest.IResponseNormalizer
}

func (v ValidationNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	if val, ok := object.(rest.Validation); ok {
		dict := map[string]interface{}{
			"FieldErrors": normalize(val.FieldErrors, options, depth),
		}

		return dict
	}

	return object
}

func (v ValidationNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.Validation)
	return
}
