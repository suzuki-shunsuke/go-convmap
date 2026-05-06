package convmap

import (
	"fmt"
	"reflect"
)

// Convert converts map[any]any to map[string]any in data.
// data isn't changed.
//
//nolint:cyclop
func Convert(data any, convertMapKey ConvertMapKey) (any, error) {
	if convertMapKey == nil {
		convertMapKey = ConvertMapKeySmart
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	//nolint:exhaustive
	switch val.Kind() {
	case reflect.Map:
		m := make(map[string]any, val.Len())
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k)
			s, err := convertMapKey(k.Interface())
			if err != nil {
				return nil, err
			}
			val, err := Convert(v.Interface(), convertMapKey)
			if err != nil {
				return nil, fmt.Errorf("key: %s: %w", s, err)
			}
			m[s] = val
		}
		return m, nil
	case reflect.Slice:
		arr := make([]any, val.Len())
		for i := 0; i < val.Len(); i++ {
			val, err := Convert(val.Index(i).Interface(), convertMapKey)
			if err != nil {
				return nil, fmt.Errorf("index: %d: %w", i, err)
			}
			arr[i] = val
		}
		return arr, nil
	default:
		return data, nil
	}
}

// ConvertMapKey converts map's key any to string.
type ConvertMapKey func(key any) (string, error)

// ConvertMapKeySmart converts any to string.
func ConvertMapKeySmart(key any) (string, error) {
	return fmt.Sprintf("%v", key), nil
}

// ConvertMapKeyStrict converts any to string.
// If key isn't a string, an error is returned.
func ConvertMapKeyStrict(key any) (string, error) {
	s, ok := key.(string)
	if !ok {
		return "", fmt.Errorf("the map key should be string: %+v", key)
	}
	return s, nil
}
