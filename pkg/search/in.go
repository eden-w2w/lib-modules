package search

import (
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"reflect"
)

var defaultOp = func(current interface{}, needle interface{}) bool {
	if current == needle {
		return true
	}
	return false
}

func In(haystack interface{}, needle interface{}, op func(current interface{}, needle interface{}) bool) (bool, int, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if op == nil {
		op = defaultOp
	}
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if op(sVal.Index(i).Interface(), needle) {
				return true, i, nil
			}
		}

		return false, -1, nil
	}

	return false, -1, general_errors.ErrUnSupportHaystack
}
