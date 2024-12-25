/*
Package convmap converts `map[interface{}]interface{}` to `map[string]interface{}`.

# Background

https://github.com/go-yaml/yaml/issues/139

When we unmarshal YAML to `interface{}`, the data type of the map will be not `map[string]interface{}` but `map[interface{}]interface{}` even if the type of all keys is string.
YAML accepts map key whose type isn't string, but JSON doesn't accept map key except for string.
And not only JSON but also some languages like Tengo (https://github.com/d5/tengo) allow only string as map key.

So this library provides the feature to convert `map[interface{}]interface{}` to `map[string]interface{}`.
*/
package convmap
