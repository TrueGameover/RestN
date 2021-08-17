package normalizer

import "RestN/rest"

type LocaleNormalizer struct {
	rest.IResponseNormalizer
}

func (n LocaleNormalizer) Normalize(object interface{}, _ rest.NormalizeMethod, _ rest.Options, _ int) interface{} {
	if locale, ok := object.(rest.Locale); ok {
		dict := map[string]interface{}{
			"Code": locale.Code,
			"Name": locale.Name,
		}

		return dict
	}

	return nil
}

func (n LocaleNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.Locale)
	return
}
