package gmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type strType string
type intType int
type uintType uint

func TestValue(t *testing.T) {
	vnil := Value{}
	vint := Value{Raw: int(3)}
	vint8 := Value{Raw: int8(3)}
	vint16 := Value{Raw: int16(3)}
	vint32 := Value{Raw: int32(3)}
	vint64 := Value{Raw: int64(3)}
	vuint := Value{Raw: uint(3)}
	vuint8 := Value{Raw: uint8(3)}
	vuint16 := Value{Raw: uint16(3)}
	vuint32 := Value{Raw: uint32(3)}
	vuint64 := Value{Raw: uint64(3)}
	vf32 := Value{Raw: float32(3)}
	vf64 := Value{Raw: float64(3)}
	vnum := Value{Raw: json.Number("3")}
	vstr := Value{Raw: string("3")}
	vbytes := Value{Raw: []byte("3")}
	vstrtp := Value{Raw: strType("3")}
	vinttp := Value{Raw: intType(3)}
	vuinttp := Value{Raw: uintType(3)}
	vbool := Value{Raw: true}
	vmap := Value{Raw: Map{"3": 3}}
	vslice := Value{Raw: Slice{3}}

	assert.True(t, vnil.IsNil())

	assert.Equal(t, int64(3), vint.IntDefaultZero())
	assert.Equal(t, int64(3), vint8.IntDefaultZero())
	assert.Equal(t, int64(3), vint16.IntDefaultZero())
	assert.Equal(t, int64(3), vint32.IntDefaultZero())
	assert.Equal(t, int64(3), vint64.IntDefaultZero())
	assert.Equal(t, int64(3), vuint.IntDefaultZero())
	assert.Equal(t, int64(3), vuint8.IntDefaultZero())
	assert.Equal(t, int64(3), vuint16.IntDefaultZero())
	assert.Equal(t, int64(3), vuint32.IntDefaultZero())
	assert.Equal(t, int64(3), vuint64.IntDefaultZero())
	assert.Equal(t, int64(3), vf32.IntDefaultZero())
	assert.Equal(t, int64(3), vf64.IntDefaultZero())
	assert.Equal(t, int64(3), vnum.IntDefaultZero())
	assert.Equal(t, int64(3), vinttp.IntDefaultZero())

	assert.Equal(t, uint64(3), vint.UintDefaultZero())
	assert.Equal(t, uint64(3), vint8.UintDefaultZero())
	assert.Equal(t, uint64(3), vint16.UintDefaultZero())
	assert.Equal(t, uint64(3), vint32.UintDefaultZero())
	assert.Equal(t, uint64(3), vint64.UintDefaultZero())
	assert.Equal(t, uint64(3), vuint.UintDefaultZero())
	assert.Equal(t, uint64(3), vuint8.UintDefaultZero())
	assert.Equal(t, uint64(3), vuint16.UintDefaultZero())
	assert.Equal(t, uint64(3), vuint32.UintDefaultZero())
	assert.Equal(t, uint64(3), vuint64.UintDefaultZero())
	assert.Equal(t, uint64(3), vf32.UintDefaultZero())
	assert.Equal(t, uint64(3), vf64.UintDefaultZero())
	assert.Equal(t, uint64(3), vnum.UintDefaultZero())
	assert.Equal(t, uint64(3), vuinttp.UintDefaultZero())

	assert.Equal(t, float64(3), vint.FloatDefaultZero())
	assert.Equal(t, float64(3), vint8.FloatDefaultZero())
	assert.Equal(t, float64(3), vint16.FloatDefaultZero())
	assert.Equal(t, float64(3), vint32.FloatDefaultZero())
	assert.Equal(t, float64(3), vint64.FloatDefaultZero())
	assert.Equal(t, float64(3), vuint.FloatDefaultZero())
	assert.Equal(t, float64(3), vuint8.FloatDefaultZero())
	assert.Equal(t, float64(3), vuint16.FloatDefaultZero())
	assert.Equal(t, float64(3), vuint32.FloatDefaultZero())
	assert.Equal(t, float64(3), vuint64.FloatDefaultZero())
	assert.Equal(t, float64(3), vf32.FloatDefaultZero())
	assert.Equal(t, float64(3), vf64.FloatDefaultZero())
	assert.Equal(t, float64(3), vnum.FloatDefaultZero())

	assert.Equal(t, string("3"), vstr.StringDefaultZero())
	assert.Equal(t, string("3"), vbytes.StringDefaultZero())
	assert.Equal(t, string("3"), vnum.StringDefaultZero())
	assert.Equal(t, string("3"), vstrtp.StringDefaultZero())

	assert.Equal(t, true, vbool.BoolDefaultZero())
	assert.Equal(t, Map{"3": 3}, vmap.MapDefaultZero())
	assert.Equal(t, Slice{3}, vslice.SliceDefaultZero())
}
