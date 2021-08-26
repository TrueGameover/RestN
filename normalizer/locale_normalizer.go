package normalizer

import (
	"github.com/TrueGameover/RestN/rest"
)

type LocaleNormalizer struct {
	rest.IResponseNormalizer
}

func (n LocaleNormalizer) Normalize(object interface{}, _ rest.NormalizeMethod, _ rest.Options, _ int) interface{} {
	locale, _ := object.(rest.Locale)

	dict := map[string]interface{}{
		"Code": locale.Code,
		"Name": locale.Name,
	}

	return &dict
}

func (n LocaleNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.Locale)
	return
}
