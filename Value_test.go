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
	vbyte := Value{Raw: byte(3)}
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
	vnum := Value{Raw: json.Number("3")}
	vstr := Value{Raw: string("3")}
	vstrtp := Value{Raw: strType("3")}
	vinttp := Value{Raw: intType(3)}
	vuinttp := Value{Raw: uintType(3)}

	assert.True(t, vnil.IsNil())
	assert.Equal(t, int64(3), vbyte.MustInt())
	assert.Equal(t, int64(3), vint.MustInt())
	assert.Equal(t, int64(3), vint8.MustInt())
	assert.Equal(t, int64(3), vint16.MustInt())
	assert.Equal(t, int64(3), vint32.MustInt())
	assert.Equal(t, int64(3), vint64.MustInt())
	assert.Equal(t, int64(3), vbyte.MustInt())
	assert.Equal(t, int64(3), vnum.MustInt())
	assert.Equal(t, int64(3), vinttp.MustInt())

	assert.Equal(t, uint64(3), vuint.MustUint())
	assert.Equal(t, uint64(3), vuint8.MustUint())
	assert.Equal(t, uint64(3), vuint16.MustUint())
	assert.Equal(t, uint64(3), vuint32.MustUint())
	assert.Equal(t, uint64(3), vuint64.MustUint())
	assert.Equal(t, uint64(3), vnum.MustUint())
	assert.Equal(t, uint64(3), vuinttp.MustUint())

	assert.Equal(t, string("3"), vstr.MustString())
	assert.Equal(t, string("3"), vstrtp.MustString())
}
