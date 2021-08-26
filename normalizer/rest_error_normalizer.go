package normalizer

import "github.com/TrueGameover/RestN/rest"

type RestErrorNormalizer struct {
	rest.IResponseNormalizer
}

func (n RestErrorNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	re, _ := object.(rest.RestError)
	dict := map[string]interface{}{
		"Message":    re.Message,
		"Validation": normalize(re.Validation, options, depth),
	}

	return &dict
}

func (n RestErrorNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.RestError)
	return
}
