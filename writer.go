package spec

import "fmt"

const WriteBufferSize = 4096

type Writer struct {
	buf  writeBuffer
	data writeData

	objects  objectStack
	elements listStack    // stack of list element tables
	fields   messageStack // stack of message field tables

	// preallocated
	_objects  [16]objectEntry
	_elements [128]listElement
	_fields   [128]messageField
}

// NewWriter returns a new writer with a default buffer.
func NewWriter() *Writer {
	buf := make([]byte, 0, WriteBufferSize)
	return NewWriterBuffer(buf)
}

// NewWriterBuffer returns a new writer with a buffer.
func NewWriterBuffer(buf []byte) *Writer {
	w := &Writer{}

	w.buf.buffer = buf[:0]
	w.data = writeData{}

	w.objects.stack = w._objects[:0]
	w.elements.stack = w._elements[:0]
	w.fields.stack = w._fields[:0]
	return w
}

// End ends writing, returns the result bytes, and resets the writer.
func (w *Writer) End() ([]byte, error) {
	if w.objects.len() > 0 {
		return nil, fmt.Errorf("end: incomplete objects, object stack size=%d", w.objects.len())
	}

	// pop data
	data := w.popData()

	// return and reset
	b := w.buf.buffer[data.start:data.end]
	w.Reset()
	return b, nil
}

// Reset clears the writer.
func (w *Writer) Reset() {
	w.buf.reset()
	w.data = writeData{}

	w.objects.reset()
	w.elements.reset()
	w.fields.reset()
}

// Primitive

func (w *Writer) Bool(v bool) error {
	start := w.buf.offset()

	if v {
		w.buf.type_(TypeTrue)
	} else {
		w.buf.type_(TypeFalse)
	}

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Byte(v byte) error {
	return w.UInt8(v)
}

func (w *Writer) Int8(v int8) error {
	start := w.buf.offset()

	w.buf.int8(v)
	w.buf.type_(TypeInt8)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Int16(v int16) error {
	start := w.buf.offset()

	w.buf.int16(v)
	w.buf.type_(TypeInt16)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Int32(v int32) error {
	start := w.buf.offset()

	w.buf.int32(v)
	w.buf.type_(TypeInt32)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Int64(v int64) error {
	start := w.buf.offset()

	w.buf.int64(v)
	w.buf.type_(TypeInt64)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) UInt8(v uint8) error {
	start := w.buf.offset()

	w.buf.uint8(v)
	w.buf.type_(TypeUInt8)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) UInt16(v uint16) error {
	start := w.buf.offset()

	w.buf.uint16(v)
	w.buf.type_(TypeUInt16)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) UInt32(v uint32) error {
	start := w.buf.offset()

	w.buf.uint32(v)
	w.buf.type_(TypeUInt32)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) UInt64(v uint64) error {
	start := w.buf.offset()

	w.buf.uint64(v)
	w.buf.type_(TypeUInt64)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Float32(v float32) error {
	start := w.buf.offset()

	w.buf.float32(v)
	w.buf.type_(TypeFloat32)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) Float64(v float64) error {
	start := w.buf.offset()

	w.buf.float64(v)
	w.buf.type_(TypeFloat64)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

// Bytes/string

func (w *Writer) Bytes(v []byte) error {
	start := w.buf.offset()

	size := w.buf.bytes(v)
	w.buf.bytesSize(size)
	w.buf.type_(TypeBytes)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

func (w *Writer) String(v string) error {
	start := w.buf.offset()

	size := w.buf.string(v)
	w.buf.stringZero()
	w.buf.stringSize(size + 1) // plus zero byte
	w.buf.type_(TypeString)

	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

// List

func (w *Writer) BeginList() error {
	// push list
	start := w.buf.offset()
	tableStart := w.elements.offset()

	w.objects.pushList(start, tableStart)
	return nil
}

func (w *Writer) Element() error {
	// pop data
	data := w.popData()
	list, err := w.objects.lastList()
	if err != nil {
		return err
	}

	// append element relative offset
	offset := uint32(data.end - list.start)
	element := listElement{offset: offset}
	w.elements.push(element)
	return nil
}

func (w *Writer) EndList() error {
	// pop list
	list, err := w.objects.popList()
	if err != nil {
		return err
	}
	dataSize := uint32(w.buf.offset() - list.start)

	// write table
	table := w.elements.pop(list.tableStart)
	tableSize := w.buf.listTable(table)

	// write sizes and type
	w.buf.listDataSize(dataSize)
	w.buf.listTableSize(tableSize)
	w.buf.type_(TypeList)

	// push data entry
	start := list.start
	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

// Message

func (w *Writer) BeginMessage() error {
	// push message
	start := w.buf.offset()
	tableStart := w.fields.offset()

	w.objects.pushMessage(start, tableStart)
	return nil
}

func (w *Writer) Field(tag uint16) error {
	// pop data
	data := w.popData()
	message, err := w.objects.lastMessage()
	if err != nil {
		return err
	}

	// insert field tag and relative offset
	f := messageField{
		tag:    tag,
		offset: uint32(data.end - message.start),
	}
	w.fields.insert(message.tableStart, f)
	return nil
}

func (w *Writer) EndMessage() error {
	// pop message
	message, err := w.objects.popMessage()
	if err != nil {
		return err
	}
	dataSize := uint32(w.buf.offset() - message.start)

	// write table
	table := w.fields.pop(message.tableStart)
	tableSize := w.buf.messageTable(table)

	// write sizes and type
	w.buf.messageDataSize(dataSize)
	w.buf.messageTableSize(tableSize)
	w.buf.type_(TypeMessage)

	// push data
	start := message.start
	end := w.buf.offset()
	w.setData(start, end)
	return nil
}

// private

// writeData holds the last written data start/end.
// There is no data stack because the data must be consumed immediatelly after it is written.
type writeData struct {
	start int
	end   int
}

func (w *Writer) setData(start, end int) {
	if w.popData().start != 0 || w.popData().end != 0 {
		panic("cannot set data, previous data not consumed")
	}

	w.data = writeData{
		start: start,
		end:   end,
	}
}

func (w *Writer) popData() writeData {
	d := w.data
	w.data = writeData{}
	return d
}
