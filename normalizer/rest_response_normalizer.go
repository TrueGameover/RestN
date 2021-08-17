package normalizer

import (
	"RestN/rest"
	"strconv"
)

type RestResponseNormalizer struct {
	rest.IResponseNormalizer
}

func (obj RestResponseNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	response, ok := object.(rest.RestResponse)

	if ok {
		dict := map[string]interface{}{}
		dict["Status"] = strconv.Itoa(response.Status)
		dict["Locale"] = normalize(response.Locale, options, depth)
		dict["Error"] = normalize(response.Error, options, depth)
		dict["Body"] = normalize(response.Body, options, depth)

		return dict
	}

	return response
}

func (obj RestResponseNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(rest.RestResponse)
	return
}
