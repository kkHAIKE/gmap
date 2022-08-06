package gmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_nextIndex(t *testing.T) {
	split := func(path string) (ret []string) {
		next := nextResult{Path: path}
		for {
			next = nextIndex(next.Path)
			ret = append(ret, next.Index)
			if !next.More {
				return
			}
		}
	}

	require.Equal(t,
		[]string{"", "foo", "ba\\r", "he.ll\\o", ""},
		split(`.foo.ba\\r.he\.ll\\o.`),
	)
}

func Test_get(t *testing.T) {
	type args struct {
		v          interface{}
		path       string
		ignoreCase bool
	}
	var zero Value
	tests := []struct {
		name string
		args args
		want Value
	}{
		{
			"normal",
			args{
				map[string]interface{}{
					"foo": []interface{}{
						map[string]interface{}{
							"ba.r": 3,
						},
					},
				},
				`foo.0.ba\.r`,
				false,
			},
			Value{true, 3},
		},
		{
			"ignore case",
			args{
				map[string]interface{}{
					"Foo": map[string]interface{}{
						"Bar": 3,
					},
				},
				`fOO.bAR`,
				true,
			},
			Value{true, 3},
		},
		{
			"out of bounds",
			args{
				[]interface{}{
					map[string]interface{}{
						"foo": 3,
					},
				},
				"1.foo",
				false,
			},
			zero,
		},
		{
			"no such key",
			args{
				map[string]interface{}{
					"foo": 3,
				},
				"bar",
				false,
			},
			zero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, get(tt.args.v, tt.args.path, tt.args.ignoreCase))
		})
	}
}
