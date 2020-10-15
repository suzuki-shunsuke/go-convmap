package convmap

import "fmt"

// Convert converts map[interface{}]interface{} to map[string]interface{} in data.
// data isn't changed.
// If the type of the key of map isn't string, an error is returned.
func Convert(data interface{}) (interface{}, error) {
	switch t := data.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(t))
		for k, v := range t {
			s, ok := k.(string)
			if !ok {
				return nil, fmt.Errorf("the map key should be string: %+v", k)
			}
			val, err := Convert(v)
			if err != nil {
				return nil, fmt.Errorf("key: %s: %w", s, err)
			}
			m[s] = val
		}
		return m, nil
	case []interface{}:
		arr := make([]interface{}, len(t))
		for i, v := range t {
			val, err := Convert(v)
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
