package pkg1

import (
	"fmt"
	"math"
	"testing"

	"github.com/complex1tech/baselibrary/types"
	"github.com/stretchr/testify/assert"
)

func TestWriteMessage__should_write_message(t *testing.T) {
	o := TestObject(t)

	m, err := o.Write(NewMessageWriter())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, m.Bool())
	assert.Equal(t, byte(255), m.Byte())

	assert.Equal(t, int32(math.MaxInt32), m.Int32())
	assert.Equal(t, int64(math.MaxInt64), m.Int64())

	assert.Equal(t, uint32(math.MaxUint32), m.Uint32())
	assert.Equal(t, uint64(math.MaxUint64), m.Uint64())

	assert.Equal(t, float32(math.MaxFloat32), m.Float32())
	assert.Equal(t, float64(math.MaxFloat64), m.Float64())

	assert.Equal(t, types.Bin64FromInt64(1), m.Bin64())
	assert.Equal(t, types.Bin128FromInt64(2), m.Bin128())
	assert.Equal(t, types.Bin256FromInt64(3), m.Bin256())

	assert.Equal(t, "hello, world", m.String())
	assert.Equal(t, []byte("goodbye, world"), m.Bytes1())

	assert.Equal(t, EnumOne, m.Enum1())
	assert.Equal(t, TestStruct(), m.Struct1())
	assert.Equal(t, "value 000", m.Submessage().Value())

	{
		list := m.Ints()
		for i := 0; i < 10; i++ {
			v := list.Get(i)
			assert.Equal(t, int64(i), v)
		}
	}

	{
		list := m.Strings()
		for i := 0; i < 10; i++ {
			v := list.Get(i)
			v0 := fmt.Sprintf("%03d", i)
			assert.Equal(t, v0, v)
		}
	}

	{
		list := m.Structs()
		for i := 0; i < 10; i++ {
			v := list.Get(i)
			v0 := Struct{
				Key:   int32(i),
				Value: -int32(i),
			}
			assert.Equal(t, v0, v)
		}
	}

	{
		list := m.Submessages()
		assert.Equal(t, 10, list.Len())
	}

	{
		list := m.Submessages1()
		assert.Equal(t, 10, list.Len())
	}
}

func TestParseMessage__should_parse_message(t *testing.T) {
	o := TestObject(t)

	m, err := o.Write(NewMessageWriter())
	if err != nil {
		t.Fatal(err)
	}
	b := m.Unwrap().Bytes()

	m1, n, err := ParseMessage(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(b), n)
	assert.Equal(t, m, m1)
}