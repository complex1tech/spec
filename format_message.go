package spec

import (
	"encoding/binary"
	"math"
)

const (
	messageFieldSize    = 1 + 2 // tag(1) + offset(2)
	messageFieldBigSize = 2 + 4 // tag(2) + offset(4)
)

// isBigList returns true if table count > uint8 or field offset > uint16.
func isBigMessage(table []messageField) bool {
	ln := len(table)
	if ln == 0 {
		return false
	}

	// count > uint8
	if ln > math.MaxUint8 {
		return true
	}

	// or offset > uint16
	last := table[ln-1]
	return last.offset > math.MaxUint16
}

// messageField specifies a field tag and a field value offset in message data array.
//
//  +----------+-------------------+
// 	| tag(1/2) |    offset(2/4)    |
//  +----------+-------------------+
//
type messageField struct {
	tag    uint16
	offset uint32
}

// messageTable is a serialized array of message fields ordered by tags.
//
//          field0                field1                field2
//  +---------------------+---------------------+---------------------+
// 	|  tag0 |   offset0   |  tag1 |   offset1   |  tag2 |   offset3   |
//  +---------------------+---------------------+---------------------+
//
type messageTable []byte

// count returns the number of fields in the table.
func (t messageTable) count(big bool) int {
	var size int
	if big {
		size = messageFieldBigSize
	} else {
		size = messageFieldSize
	}
	return len(t) / size
}

// field returns a field by its index or false,
func (t messageTable) field(big bool, i int) (f messageField, ok bool) {
	// inline size
	var size int
	if big {
		size = messageFieldBigSize
	} else {
		size = messageFieldSize
	}

	// count
	n := len(t) / size
	switch {
	case i < 0:
		return
	case i >= n:
		return
	}

	off := i * size
	b := t[off : off+size]

	if big {
		f = messageField{
			tag:    binary.BigEndian.Uint16(b),
			offset: binary.BigEndian.Uint32(b[2:]),
		}
	} else {
		f = messageField{
			tag:    uint16(b[0]),
			offset: uint32(binary.BigEndian.Uint16(b[1:])),
		}
	}

	ok = true
	return
}

// fields parses the table and returns a slice of fields.
func (t messageTable) fields(big bool) []messageField {
	n := t.count(big)

	result := make([]messageField, 0, n)
	for i := 0; i < n; i++ {
		field, ok := t.field(big, i)
		if !ok {
			continue
		}
		result = append(result, field)
	}
	return result
}

// offset

// offset returns field start/end by its tag or -1/-1.
func (t messageTable) offset(big bool, tag uint16) (int, int) {
	if big {
		return t._offset_big(tag)
	} else {
		return t._offset_small(tag)
	}
}

func (t messageTable) _offset_big(tag uint16) (int, int) {
	size := messageFieldBigSize
	n := len(t) / size

	// binary search table
	left, right := 0, (n - 1)
	for left <= right {
		// middle
		middle := int(uint(left+right) >> 1) // avoid overflow

		// offset
		off := middle * size
		b := t[off : off+size]

		// current tag
		cur := binary.BigEndian.Uint16(b)

		// check current
		switch {
		case cur < tag:
			left = middle + 1

		case cur > tag:
			right = middle - 1

		case cur == tag:
			// start
			var start int
			if middle > 0 {
				start = int(binary.BigEndian.Uint32(t[off-4:]))
			}

			// end
			end := int(binary.BigEndian.Uint32(b[2:]))
			return start, end
		}
	}

	return -1, -1
}

func (t messageTable) _offset_small(tag uint16) (int, int) {
	size := messageFieldSize
	n := len(t) / size

	// binary search table
	left, right := 0, (n - 1)
	for left <= right {
		// middle
		middle := int(uint(left+right) >> 1) // avoid overflow

		// offset
		off := middle * size
		b := t[off : off+size]

		// current tag
		cur := uint16(b[0])

		// check current
		switch {
		case cur < tag:
			left = middle + 1

		case cur > tag:
			right = middle - 1

		case cur == tag:
			// start
			var start int
			if middle > 0 {
				start = int(binary.BigEndian.Uint16(t[off-2:]))
			}

			// end
			end := int(binary.BigEndian.Uint16(b[1:]))
			return start, end
		}
	}

	return -1, -1
}

// offsetByIndex

// offsetByIndex returns field start/end by its index or -1/-1.
func (t messageTable) offsetByIndex(big bool, i int) (int, int) {
	if big {
		return t._offsetByIndex_big(i)
	} else {
		return t._offsetByIndex_small(i)
	}
}

func (t messageTable) _offsetByIndex_big(i int) (int, int) {
	size := messageFieldBigSize
	n := len(t) / size

	// check count
	switch {
	case i < 0:
		return -1, -1
	case i >= n:
		return -1, -1
	}

	// offset
	off := i * size

	// start
	var start int
	if i > 0 {
		start = int(binary.BigEndian.Uint32(t[off-4:]))
	}

	// end
	end := int(binary.BigEndian.Uint32(t[off+2:]))
	return start, end
}

func (t messageTable) _offsetByIndex_small(i int) (int, int) {
	size := messageFieldSize
	n := len(t) / size

	// check count
	switch {
	case i < 0:
		return -1, -1
	case i >= n:
		return -1, -1
	}

	// offset
	off := i * size

	// start
	var start int
	if i > 0 {
		start = int(binary.BigEndian.Uint16(t[off-2:]))
	}

	// end
	end := int(binary.BigEndian.Uint16(t[off+1:]))
	return start, end
}