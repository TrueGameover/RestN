package rest

type NormalizeMethod func(obj interface{}, options Options, depth int) interface{}

type IResponseNormalizer interface {
	Normalize(object interface{}, normalize NormalizeMethod, options Options, depth int) interface{}
	Support(object interface{}) bool
}

const DefaultDepth = 100

var normalizers []IResponseNormalizer

func RegisterNormalizer(normalizer IResponseNormalizer) {
	normalizers = append(normalizers, normalizer)
}

func ClearAllNormalizers() {
	normalizers = nil
}

func (response *RestResponse) NormalizeResponse() interface{} {
	if response.depth == 0 {
		response.depth = DefaultDepth
	}

	return normalize(*response, response.options, response.depth)
}

func normalize(obj interface{}, options Options, depth int) interface{} {
	depth--
	if depth < 0 {
		return nil
	}

	for _, n := range normalizers {
		if n.Support(obj) {
			return n.Normalize(obj, normalize, options, depth)
		}
	}

	return obj
}
