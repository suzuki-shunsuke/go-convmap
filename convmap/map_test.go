package convmap_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-convmap/convmap"
)

func TestConvert(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		title         string
		isErr         bool
		input         interface{}
		exp           interface{}
		convertMapKey convmap.ConvertMapKey
	}{
		{
			title: "int",
			input: 1,
			exp:   1,
		},
		{
			title: "string",
			input: "hello",
			exp:   "hello",
		},
		{
			title: "map[interface{}]interface{}",
			input: map[interface{}]interface{}{
				"foo": "bar",
			},
			exp: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			title: "[]map[interface{}]interface{}",
			input: []interface{}{
				map[interface{}]interface{}{
					"foo": "bar",
				},
			},
			exp: []interface{}{
				map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		{
			title: "key should be string",
			isErr: true,
			input: []interface{}{
				map[interface{}]interface{}{
					true: "bar",
				},
			},
			convertMapKey: convmap.ConvertMapKeyStrict,
		},
		{
			title: "failed to convert map's value",
			isErr: true,
			input: map[interface{}]interface{}{
				"foo": map[interface{}]interface{}{
					true: "bar",
				},
			},
			convertMapKey: convmap.ConvertMapKeyStrict,
		},
		{
			title: "interface is nested",
			input: map[string]interface{}{
				"foo": []interface{}{
					map[interface{}]interface{}{
						"key":      "sometype",
						"operator": "In",
						"values": []interface{}{
							"large",
							"medium",
						},
						"extra": map[interface{}]interface{}{
							"comple": map[interface{}]interface{}{
								"ca": "ted",
							},
						},
					},
				},
			},
			exp: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"key":      "sometype",
						"operator": "In",
						"values": []interface{}{
							"large",
							"medium",
						},
						"extra": map[string]interface{}{
							"comple": map[string]interface{}{
								"ca": "ted",
							},
						},
					},
				},
			},
		},
		{
			title: "special type map[uint16]interface{}",
			input: map[uint16]interface{}{
				1234: "bar",
			},
			exp: map[string]interface{}{
				"1234": "bar",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			a, err := convmap.Convert(d.input, d.convertMapKey)
			if d.isErr {
				if err == nil {
					t.Error("convmap.Convert() should return an error")
				}
				return
			}
			if err != nil {
				t.Error(err)
				return
			}
			if d.exp == nil {
				return
			}
			if !reflect.DeepEqual(d.exp, a) {
				t.Errorf("want %+v, got %+v", d.exp, a)
			}
		})
	}
}
