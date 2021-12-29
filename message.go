package spec

type Message struct {
	data  []byte
	table messageTable

	body uint32 // body size
	big  bool   // big/small table format
}

// GetMessage parses and returns a message, but does not validate it.
func GetMessage(b []byte) (Message, error) {
	return readMessage(b)
}

// ReadMessage reads, recursively validates and returns a message.
func ReadMessage(b []byte) (Message, error) {
	m, err := readMessage(b)
	if err != nil {
		return Message{}, err
	}
	if err := m.Validate(); err != nil {
		return Message{}, err
	}
	return m, nil
}

// Data returns the exact message bytes.
func (m Message) Data() []byte {
	return m.data
}

// Validate recursively validates the message.
func (m Message) Validate() error {
	n := m.Len()

	for i := 0; i < n; i++ {
		data := m.FieldByIndex(i)
		if len(data) == 0 {
			continue
		}
		if _, err := ReadValue(data); err != nil {
			return err
		}
	}
	return nil
}

// Field returns a field data by a tag or nil.
func (m Message) Field(tag uint16) []byte {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return nil
	case end > int(m.body):
		return nil
	}
	return m.data[start:end]
}

// FieldByIndex returns a field data by an index or nil.
func (m Message) FieldByIndex(i int) []byte {
	start, end := m.table.offsetByIndex(m.big, i)
	switch {
	case start < 0:
		return nil
	case end > int(m.body):
		return nil
	}
	return m.data[start:end]
}

// Len returns the number of fields in the message.
func (m Message) Len() int {
	return m.table.count(m.big)
}

// Getters

func (m Message) Bool(tag uint16) bool {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return false
	case end > int(m.body):
		return false
	}

	b := m.data[start:end]
	v, _ := ReadBool(b)
	return v
}

func (m Message) Int8(tag uint16) int8 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadInt8(b)
	return v
}

func (m Message) Int16(tag uint16) int16 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadInt16(b)
	return v
}

func (m Message) Int32(tag uint16) int32 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadInt32(b)
	return v
}

func (m Message) Int64(tag uint16) int64 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadInt64(b)
	return v
}

func (m Message) Uint8(tag uint16) uint8 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadUint8(b)
	return v
}

func (m Message) Uint16(tag uint16) uint16 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadUint16(b)
	return v
}

func (m Message) Uint32(tag uint16) uint32 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadUint32(b)
	return v
}

func (m Message) Uint64(tag uint16) uint64 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadUint64(b)
	return v
}

func (m Message) Float32(tag uint16) float32 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadFloat32(b)
	return v
}

func (m Message) Float64(tag uint16) float64 {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return 0
	case end > int(m.body):
		return 0
	}

	b := m.data[start:end]
	v, _ := ReadFloat64(b)
	return v
}

func (m Message) Bytes(tag uint16) []byte {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return nil
	case end > int(m.body):
		return nil
	}

	b := m.data[start:end]
	v, _ := readBytes(b)
	return v
}

func (m Message) String(tag uint16) string {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return ""
	case end > int(m.body):
		return ""
	}

	b := m.data[start:end]
	v, _ := readString(b)
	return v
}

func (m Message) List(tag uint16) List {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return List{}
	case end > int(m.body):
		return List{}
	}

	b := m.data[start:end]
	v, _ := GetList(b)
	return v
}

func (m Message) Message(tag uint16) Message {
	start, end := m.table.offset(m.big, tag)
	switch {
	case start < 0:
		return Message{}
	case end > int(m.body):
		return Message{}
	}

	b := m.data[start:end]
	v, _ := GetMessage(b)
	return v
}
