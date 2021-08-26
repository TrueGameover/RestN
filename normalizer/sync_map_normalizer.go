package normalizer

import (
	"github.com/TrueGameover/RestN/rest"
	"sync"
)

type SyncMapNormalizer struct {
	rest.IResponseNormalizer
}

func (v SyncMapNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	val, _ := object.(*sync.Map)
	dict := map[string]interface{}{}

	val.Range(func(key, value interface{}) bool {
		if str, ok := key.(string); ok {
			dict[str] = value
		}

		return true
	})

	return dict
}

func (v SyncMapNormalizer) Support(object interface{}) (ok bool) {
	_, ok = object.(*sync.Map)
	return
}
