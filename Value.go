package gmap

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type Value struct {
	// check exists
	Has bool
	Raw interface{}
}

func (v Value) IsNil() bool {
	return v.Raw == nil
}

func (v Value) Int() (int64, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return 0, false
	case int:
		return int64(t), true
	case int8:
		return int64(t), true
	case int16:
		return int64(t), true
	case int32:
		return int64(t), true
	case int64:
		return t, true
	case uint:
		return int64(t), true
	case uint8:
		return int64(t), true
	case uint16:
		return int64(t), true
	case uint32:
		return int64(t), true
	case uint64:
		if t > math.MaxInt64 {
			return 0, false
		}
		return int64(t), true
	case float32:
		if r := int64(t); float32(r) == t {
			return r, true
		}
		return 0, false
	case float64:
		if r := int64(t); float64(r) == t {
			return r, true
		}
		return 0, false
	case json.Number:
		if r, err := t.Int64(); err == nil {
			return r, true
		}
		return 0, false
	}

	rv := reflect.ValueOf(v.Raw)
	if rv.CanInt() {
		return rv.Int(), true
	}

	return 0, false
}

func (v Value) MustInt() int64 {
	r, ok := v.Int()
	if !ok {
		panic(fmt.Sprintf("not int: %v", v.Raw))
	}
	return r
}

func (v Value) Uint() (uint64, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return 0, false
	case int:
		if t < 0 {
			return 0, false
		}
		return uint64(t), true
	case int8:
		if t < 0 {
			return 0, false
		}
		return uint64(t), true
	case int16:
		if t < 0 {
			return 0, false
		}
		return uint64(t), true
	case int32:
		if t < 0 {
			return 0, false
		}
		return uint64(t), true
	case int64:
		if t < 0 {
			return 0, false
		}
		return uint64(t), true
	case uint:
		return uint64(t), true
	case uint8:
		return uint64(t), true
	case uint16:
		return uint64(t), true
	case uint32:
		return uint64(t), true
	case uint64:
		return t, true
	case float32:
		if r := uint64(t); float32(r) == t {
			return r, true
		}
		return 0, false
	case float64:
		if r := uint64(t); float64(r) == t {
			return r, true
		}
		return 0, false
	case json.Number:
		if r, err := strconv.ParseUint(string(t), 10, 64); err == nil {
			return r, true
		}
		return 0, false
	}

	rv := reflect.ValueOf(v.Raw)
	if rv.CanUint() {
		return rv.Uint(), true
	}

	return 0, false
}

func (v Value) MustUint() uint64 {
	r, ok := v.Uint()
	if !ok {
		panic(fmt.Sprintf("not uint: %v", v.Raw))
	}
	return r
}

func (v Value) Float() (float64, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return 0, false
	case int:
		return float64(t), true
	case int8:
		return float64(t), true
	case int16:
		return float64(t), true
	case int32:
		return float64(t), true
	case int64:
		if r := float64(t); int64(r) == t {
			return r, true
		}
		return 0, false
	case uint:
		return float64(t), true
	case uint8:
		return float64(t), true
	case uint16:
		return float64(t), true
	case uint32:
		return float64(t), true
	case uint64:
		if r := float64(t); uint64(r) == t {
			return r, true
		}
		return 0, false
	case float32:
		return float64(t), true
	case float64:
		return t, true
	case json.Number:
		if r, err := t.Float64(); err == nil {
			return r, true
		}
		return 0, false
	}

	return 0, false
}

func (v Value) MustFloat() float64 {
	r, ok := v.Float()
	if !ok {
		panic(fmt.Sprintf("not float: %v", v.Raw))
	}
	return r
}

func (v Value) String() (string, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return "", false
	case string:
		return t, true
	case []byte:
		return string(t), true
	case fmt.Stringer:
		return t.String(), true
	}

	rv := reflect.ValueOf(v.Raw)
	if rv.Kind() == reflect.String {
		return rv.String(), true
	}

	return "", false
}

func (v Value) MustString() string {
	r, ok := v.String()
	if !ok {
		panic(fmt.Sprintf("not string: %v", v.Raw))
	}
	return r
}

func (v Value) Bool() (bool, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return false, false
	case bool:
		return t, true
	}

	return false, false
}

func (v Value) MustBool() bool {
	r, ok := v.Bool()
	if !ok {
		panic(fmt.Sprintf("not bool: %v", v.Raw))
	}
	return r
}

func (v Value) Slice() (Slice, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return nil, false
	case Slice:
		return t, true
	case []interface{}:
		return t, true
	}

	return nil, false
}

func (v Value) MustSlice() Slice {
	r, ok := v.Slice()
	if !ok {
		panic(fmt.Sprintf("not slice: %v", v.Raw))
	}
	return r
}

func (v Value) Map() (Map, bool) {
	switch t := v.Raw.(type) {
	case nil:
		return nil, false
	case Map:
		return t, true
	case map[string]interface{}:
		return t, true
	}

	return nil, false
}

func (v Value) MustMap() Map {
	r, ok := v.Map()
	if !ok {
		panic(fmt.Sprintf("not map: %v", v.Raw))
	}
	return r
}
