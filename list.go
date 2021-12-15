package protocol

import (
	"encoding/binary"
	"fmt"
)

type List struct {
	bytes []byte

	type_     Type
	tableSize uint32
	dataSize  uint32
	table     listTable
	data      readBuffer
}

func ReadList(p []byte) List {
	buf := readBuffer(p)
	type_, b := buf.type_()
	if type_ != TypeList {
		return List{}
	}

	tableSize, b := b.listTableSize()
	dataSize, b := b.listDataSize()
	table, b := b.listTable(tableSize)
	data, _ := b.listData(dataSize)
	bytes, _ := buf.listBytes(tableSize, dataSize) // slice initial buffer

	return List{
		bytes: bytes,

		type_:     type_,
		tableSize: tableSize,
		dataSize:  dataSize,
		table:     table,
		data:      data,
	}
}

// Element returns an element value by an index or false.
func (l List) Element(i int) (reader, bool) {
	elem, ok := l.table.lookup(i)
	if !ok {
		return reader{}, false
	}

	buf := l.data.listElement(elem.offset)
	return read(buf), true
}

// Elements returns the number of elements in the list.
func (l List) Elements() int {
	return l.table.count()
}

// listElement specifies an element value offset in list data array.
//
//  +-----------------+
// 	|    offset(4)    |
//  +-----------------+
//
type listElement struct {
	offset uint32
}

const listElementSize = 4

// listTable is a serialized array of list element offsets ordered by index.
//
//  +--------------+--------------+--------------+
// 	|    off0(4)   |    off1(4)   |    off2(4)   |
//  +--------------+--------------+--------------+
//
type listTable []byte

// readListTable casts bytes into an element table,
// returns an error if length is not divisible by listElementSize.
func readListTable(data []byte) (listTable, error) {
	ln := len(data)
	if (ln % listElementSize) != 0 {
		return nil, fmt.Errorf(
			"read element table: invalid table length, must be divisible by %d, length=%v",
			listElementSize, ln,
		)
	}

	return data, nil
}

// writeListTable writes elements to a binary element table.
// used in tests.
func writeListTable(elements []listElement) listTable {
	// alloc table
	size := len(elements) * listElementSize
	result := make([]byte, size)

	// write elements
	for i, elem := range elements {
		off := i * listElementSize
		b := result[off:]

		binary.BigEndian.PutUint32(b, uint32(elem.offset))
	}

	return result
}

// get returns an element by its index, panics if index is out of range.
func (t listTable) get(i int) listElement {
	n := t.count()
	if i >= n {
		panic(fmt.Sprintf("get element: index out of range, length=%d, index=%d", n, i))
	}

	off := i * listElementSize
	b := t[off : off+listElementSize]
	elem := listElement{offset: binary.BigEndian.Uint32(b)}
	return elem
}

// lookup returns an element by its index or false.
func (t listTable) lookup(i int) (listElement, bool) {
	n := t.count()
	if i >= n {
		return listElement{}, false
	}

	off := i * listElementSize
	b := t[off : off+listElementSize]
	elem := listElement{offset: binary.BigEndian.Uint32(b)}
	return elem, true
}

// count returns the number of elements in the table.
func (t listTable) count() int {
	return len(t) / listElementSize
}

// elements parses the table and returns a slice of elements
func (t listTable) elements() []listElement {
	n := t.count()

	result := make([]listElement, 0, n)
	for i := 0; i < n; i++ {
		elem := t.get(i)
		result = append(result, elem)
	}
	return result
}
