package convmap

import (
	"fmt"
	"reflect"
)

// Convert converts map[interface{}]interface{} to map[string]interface{} in data.
// data isn't changed.
//nolint:cyclop
func Convert(data interface{}, convertMapKey ConvertMapKey) (interface{}, error) {
	if convertMapKey == nil {
		convertMapKey = ConvertMapKeySmart
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	//nolint:exhaustive
	switch val.Kind() {
	case reflect.Map:
		m := make(map[string]interface{}, val.Len())
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
		arr := make([]interface{}, val.Len())
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

// ConvertMapKey converts map's key interface{} to string.
type ConvertMapKey func(key interface{}) (string, error)

// ConvertMapKeySmart converts interface{} to string.
func ConvertMapKeySmart(key interface{}) (string, error) {
	return fmt.Sprintf("%v", key), nil
}

// ConvertMapKeyStrict converts interface{} to string.
// If key isn't a string, an error is returned.
func ConvertMapKeyStrict(key interface{}) (string, error) {
	s, ok := key.(string)
	if !ok {
		return "", fmt.Errorf("the map key should be string: %+v", key)
	}
	return s, nil
}
