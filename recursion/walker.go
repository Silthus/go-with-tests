package recursion

import (
	"reflect"
)

func Walk(object any, fn func(field string)) {
	var val = reflect.ValueOf(object)
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Pointer:
		Walk(val.Elem().Interface(), fn)
	case reflect.Chan:
		walkChannel(val, fn)
	case reflect.Slice, reflect.Array:
		walkSlice(val, fn)
	case reflect.Map:
		walkMap(val, fn)
	case reflect.Struct:
		walkFields(val, fn)
	case reflect.Func:
		walkFunction(val, fn)
	}
}

func walkFunction(val reflect.Value, fn func(field string)) {
	fnResult := val.Call(nil)
	for _, res := range fnResult {
		walkValue(res, fn)
	}
}

func walkChannel(val reflect.Value, fn func(field string)) {
	for v, ok := val.Recv(); ok; v, ok = val.Recv() {
		walkValue(v, fn)
	}
}

func walkMap(val reflect.Value, fn func(field string)) {
	for _, key := range val.MapKeys() {
		walkValue(val.MapIndex(key), fn)
	}
}

func walkSlice(val reflect.Value, fn func(field string)) {
	for i := 0; i < val.Len(); i++ {
		walkValue(val.Index(i), fn)
	}
}

func walkFields(val reflect.Value, fn func(field string)) {
	for i := 0; i < val.NumField(); i++ {
		walkValue(val.Field(i), fn)
	}
}

func walkValue(field reflect.Value, fn func(field string)) {
	switch field.Kind() {
	case reflect.String:
		fn(field.String())
	case reflect.Struct:
		Walk(field.Interface(), fn)
	case reflect.Slice:
		walkSlice(field, fn)
	}
}
