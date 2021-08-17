package normalizer

import (
	"RestN/rest"
	"reflect"
)

type SliceNormalizer struct {
	rest.IResponseNormalizer
}

func (s SliceNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
	rg := reflect.ValueOf(object)
	var arr []interface{}

	for i := 0; i < rg.Len(); i++ {
		value := rg.Index(i).Interface()
		arr = append(arr, normalize(value, options, depth))
	}

	return arr
}

func (s SliceNormalizer) Support(object interface{}) (ok bool) {
	if object != nil {
		t := reflect.TypeOf(object)
		kind := t.Kind()

		if kind == reflect.Slice || kind == reflect.Array {
			ok = true
		}
	}

	return
}
