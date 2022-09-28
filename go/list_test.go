package spec

import (
	"testing"
)

func testEncodeList(t *testing.T) []byte {
	w := NewEncoder()
	w.BeginList()

	w.Int32(1)
	w.Element()
	w.Int64(1)
	w.Element()

	w.Uint32(1)
	w.Element()
	w.Uint64(1)
	w.Element()

	w.Float32(1)
	w.Element()
	w.Float64(1)
	w.Element()

	w.String("hello, world")
	w.Element()
	w.Bytes([]byte("hello, world"))
	w.Element()

	w.BeginList()
	w.String("element1")
	w.Element()
	w.End()
	w.Element()

	w.BeginMessage()
	w.String("field1")
	w.Field(1)
	w.End()
	w.Element()

	b, err := w.End()
	if err != nil {
		t.Fatal(err)
	}
	return b
}