# go-convmap

[![Build Status](https://github.com/suzuki-shunsuke/go-convmap/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/go-convmap/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/go-convmap)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-convmap/convmap)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b1b3bbcbc72ad25f6c44/test_coverage)](https://codeclimate.com/github/suzuki-shunsuke/go-convmap/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-convmap)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-convmap)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-convmap.svg)](https://github.com/suzuki-shunsuke/go-convmap)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-convmap/master/LICENSE)

Go library to convert `map[interface{}]interface{}` to `map[string]interface{}`.

## Background

https://github.com/go-yaml/yaml/issues/139

When we unmarshal YAML to `interface{}`, the data type of the map will be not `map[string]interface{}` but `map[interface{}]interface{}` even if the type of all keys is string.
YAML accepts map key whose type isn't string, but JSON doesn't accept map key except for string.
And not only JSON but also some languages like [Tengo](https://github.com/d5/tengo) allow only string as map key.

So this library provides the feature to convert `map[interface{}]interface{}` to `map[string]interface{}`.

## LICENSE

[MIT](LICENSE)
