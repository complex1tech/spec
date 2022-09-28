package spec

import (
	"math"
	"testing"

	"github.com/complex1tech/baselibrary/buffer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// DecodeType

func TestDecodeType__should_return_type(t *testing.T) {
	b := []byte{}
	b = append(b, byte(TypeString))

	v, n, err := DecodeType(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, len(b))
	assert.Equal(t, v, TypeString)
}

func TestDecodeType__should_return_undefined_when_empty(t *testing.T) {
	b := []byte{}

	v, n, err := DecodeType(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Zero(t, n)
	assert.Equal(t, v, TypeUndefined)
}

// DecodeBool

func TestDecodeBool__should_decode_bool_value(t *testing.T) {
	b := []byte{byte(TypeTrue)}
	v, n, err := DecodeBool(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, len(b))
	assert.Equal(t, true, v)

	b = []byte{byte(TypeFalse)}
	v, n, err = DecodeBool(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, len(b))
	assert.Equal(t, false, v)
}

// DecodeByte

func TestDecodeByte__should_decode_byte(t *testing.T) {
	b := buffer.New()
	EncodeByte(b, 1)

	v, n, err := DecodeByte(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, byte(1), v)
}

// Int32

func TestDecodeInt32__should_decode_int32(t *testing.T) {
	b := buffer.New()
	EncodeInt32(b, 1)

	v, n, err := DecodeInt32(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, int32(1), v)
}

// Int64

func TestDecodeInt64__should_decode_int64(t *testing.T) {
	b := buffer.New()
	EncodeInt64(b, 1)

	v, n, err := DecodeInt64(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, int64(1), v)
}

func TestDecodeInt64__should_decode_int64_from_int32(t *testing.T) {
	b := buffer.New()
	EncodeInt32(b, 1)

	v, n, err := DecodeInt64(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, int64(1), v)
}

// Float32

func TestDecodeFloat32__should_decode_float32(t *testing.T) {
	b := buffer.New()
	EncodeFloat32(b, math.MaxFloat32)

	v, n, err := DecodeFloat32(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, float32(math.MaxFloat32), v)
}

func TestDecodeFloat32__should_decode_float32_from_float64(t *testing.T) {
	b := buffer.New()
	EncodeFloat64(b, math.MaxFloat32)

	v, n, err := DecodeFloat32(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, float32(math.MaxFloat32), v)
}

// Float64

func TestDecodeFloat64__should_decode_float64(t *testing.T) {
	b := buffer.New()
	EncodeFloat64(b, math.MaxFloat64)

	v, n, err := DecodeFloat64(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, float64(math.MaxFloat64), v)
}

func TestDecodeFloat64__should_decode_float64_from_float32(t *testing.T) {
	b := buffer.New()
	EncodeFloat32(b, math.MaxFloat32)

	v, n, err := DecodeFloat64(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, float64(math.MaxFloat32), v)
}

// Bytes

func TestDecodeBytes__should_decode_bytes(t *testing.T) {
	v := []byte("hello, world")

	b := buffer.New()
	_, err := EncodeBytes(b, v)
	if err != nil {
		t.Fatal(err)
	}

	v1, n, err := DecodeBytes(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, v, v1)
}

// String

func TestDecodeString__should_decode_string(t *testing.T) {
	v := "hello, world"

	b := buffer.New()
	_, err := EncodeString(b, v)
	if err != nil {
		t.Fatal(err)
	}

	v1, n, err := DecodeString(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, b.Len())
	assert.Equal(t, v, v1)
}

// List

func TestDecodeListMeta__should_decode_list(t *testing.T) {
	b := testEncodeList(t)
	_, n, err := decodeListMeta(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, len(b))
}

func TestDecodeListTable__should_decode_list_table(t *testing.T) {
	elements := testListElements()

	for i := 0; i <= len(elements); i++ {
		b := buffer.New()
		ee0 := elements[i:]

		size, err := encodeListTable(b, ee0, false)
		if err != nil {
			t.Fatal(err)
		}

		table1, err := decodeListTable(b.Bytes(), uint32(size), false)
		if err != nil {
			t.Fatal(err)
		}

		ee1 := table1.elements(false)
		require.Equal(t, ee0, ee1)
	}
}

func TestDecodeListMeta__should_return_error_when_invalid_type(t *testing.T) {
	b := testEncodeList(t)
	b[len(b)-1] = byte(TypeMessage)

	_, _, err := decodeListMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid type")
}

func TestDecodeListMeta__should_return_error_when_invalid_table_size(t *testing.T) {
	b := []byte{}
	b = append(b, 0xff)
	b = append(b, byte(TypeList))

	_, _, err := decodeListMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid table size")
}

func TestDecodeListMeta__should_return_error_when_invalid_data_size(t *testing.T) {
	big := false
	b := []byte{}
	b = append(b, 0xff)
	b = appendSize(b, big, 1000)
	b = append(b, byte(TypeList))

	_, _, err := decodeListMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid data size")
}

func TestDecodeListMeta__should_return_error_when_invalid_table(t *testing.T) {
	buf := buffer.New()
	_, err := encodeListTable(buf, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	big := false
	b := buf.Bytes()
	b = appendSize(b, big, 0)    // data size
	b = appendSize(b, big, 1000) // table size
	b = append(b, byte(TypeList))

	_, _, err = decodeListMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid table")
}

func TestDecodeListMeta__should_return_error_when_invalid_data(t *testing.T) {
	buf := buffer.New()
	_, err := encodeListTable(buf, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	big := false
	b := buf.Bytes()
	b = appendSize(b, big, 1000) // data size
	b = appendSize(b, big, 0)    // table size
	b = append(b, byte(TypeList))

	_, _, err = decodeListMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid data")
}

// Message

func TestDecodeMessageMeta__should_decode_message_meta(t *testing.T) {
	b := testEncodeMessage(t)
	_, n, err := decodeMessageMeta(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, len(b))
}

func TestDecodeMessageTable__should_decode_message_table(t *testing.T) {
	fields := testMessageFields()

	for i := 0; i <= len(fields); i++ {
		buf := buffer.New()
		fields0 := fields[i:]

		size, err := encodeMessageTable(buf, fields0, false)
		if err != nil {
			t.Fatal(err)
		}

		table1, err := decodeMessageTable(buf.Bytes(), uint32(size), false)
		if err != nil {
			t.Fatal(err)
		}

		fields1 := table1.fields(false)
		require.Equal(t, fields0, fields1)
	}
}

func TestDecodeMessageMeta__should_return_error_when_invalid_type(t *testing.T) {
	b := testEncodeMessage(t)
	b[len(b)-1] = byte(TypeList)

	_, _, err := decodeMessageMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid type")
}

func TestDecodeMessageMeta__should_return_error_when_invalid_table_size(t *testing.T) {
	b := []byte{}
	b = append(b, 0xff)
	b = append(b, byte(TypeMessage))

	_, _, err := decodeMessageMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid table size")
}

func TestDecodeMessageMeta__should_return_error_when_invalid_data_size(t *testing.T) {
	big := false
	b := []byte{}
	b = append(b, 0xff)
	b = appendSize(b, big, 1000)
	b = append(b, byte(TypeMessage))

	_, _, err := decodeMessageMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid data size")
}

func TestDecodeMessageMeta__should_return_error_when_invalid_table(t *testing.T) {
	buf := buffer.New()
	_, err := encodeMessageTable(buf, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	big := false
	b := buf.Bytes()
	b = appendSize(b, big, 0)    // data size
	b = appendSize(b, big, 1000) // table size
	b = append(b, byte(TypeMessage))

	_, _, err = decodeMessageMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid table")
}

func TestDecodeMessageMeta__should_return_error_when_invalid_data(t *testing.T) {
	buf := buffer.New()

	_, err := encodeMessageTable(buf, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	big := false
	b := buf.Bytes()
	b = appendSize(b, big, 1000)
	b = appendSize(b, big, 0)
	b = append(b, byte(TypeMessage))

	_, _, err = decodeMessageMeta(b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid data")
}