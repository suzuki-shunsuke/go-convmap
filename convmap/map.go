package convmap

import (
	"fmt"
	"strconv"
)

// Convert converts map[interface{}]interface{} to map[string]interface{} in data.
// data isn't changed.
func Convert(data interface{}, convertMapKey ConvertMapKey) (interface{}, error) {
	if convertMapKey == nil {
		convertMapKey = ConvertMapKeySmart
	}
	switch t := data.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(t))
		for k, v := range t {
			s, err := convertMapKey(k)
			if err != nil {
				return nil, err
			}
			val, err := Convert(v, convertMapKey)
			if err != nil {
				return nil, fmt.Errorf("key: %s: %w", s, err)
			}
			m[s] = val
		}
		return m, nil
	case []interface{}:
		arr := make([]interface{}, len(t))
		for i, v := range t {
			val, err := Convert(v, convertMapKey)
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
// If key is bool, int, or int64, it is converted to string with strconv package.
func ConvertMapKeySmart(key interface{}) (string, error) {
	switch k := key.(type) {
	case string:
		return k, nil
	case bool:
		return strconv.FormatBool(k), nil
	case int:
		return strconv.Itoa(k), nil
	case int64:
		return strconv.FormatInt(k, 10), nil //nolint:gomnd
	default:
		return "", fmt.Errorf("the map key should be string: %+v", key)
	}
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
