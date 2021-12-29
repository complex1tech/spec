package spec

type Value []byte

// GetValue parses and returns a value, but does not validate it.
func GetValue(b []byte) (Value, error) {
	t, err := ReadType(b)
	if err != nil {
		return Value{}, err
	}
	if err := CheckType(t); err != nil {
		return Value{}, err
	}
	return Value(b), nil
}

// ReadValue reads, recursively validates and returns a value.
func ReadValue(b []byte) (Value, error) {
	v := Value(b)
	if err := v.Validate(); err != nil {
		return Value{}, err
	}
	return v, nil
}

func (v Value) Validate() error {
	t, err := ReadType(v)
	if err != nil {
		return err
	}

	switch t {
	case TypeNil, TypeTrue, TypeFalse:
		return nil

	case TypeInt8:
		_, err = ReadInt8(v)
	case TypeInt16:
		_, err = ReadInt16(v)
	case TypeInt32:
		_, err = ReadInt32(v)
	case TypeInt64:
		_, err = ReadInt64(v)

	case TypeUint8:
		_, err = ReadUint8(v)
	case TypeUint16:
		_, err = ReadUint16(v)
	case TypeUint32:
		_, err = ReadUint32(v)
	case TypeUint64:
		_, err = ReadUint64(v)

	case TypeFloat32:
		_, err = ReadFloat32(v)
	case TypeFloat64:
		_, err = ReadFloat64(v)

	case TypeBytes:
		_, err = ReadBytes(v)
	case TypeString:
		_, err = ReadString(v)

	case TypeList:
		_, err = ReadList(v)
	case TypeMessage:
		_, err = ReadMessage(v)
	}
	return err
}

func (v Value) Type() Type {
	p, _ := ReadType(v)
	return p
}

func (v Value) Nil() bool {
	p, _ := ReadBool(v)
	return p
}

func (v Value) Bool() bool {
	p, _ := ReadBool(v)
	return p
}

func (v Value) Byte() byte {
	p, _ := ReadByte(v)
	return p
}

func (v Value) Int8() int8 {
	p, _ := ReadInt8(v)
	return p
}

func (v Value) Int16() int16 {
	p, _ := ReadInt16(v)
	return p
}

func (v Value) Int32() int32 {
	p, _ := ReadInt32(v)
	return p
}

func (v Value) Int64() int64 {
	p, _ := ReadInt64(v)
	return p
}

func (v Value) Uint8() uint8 {
	p, _ := ReadUint8(v)
	return p
}

func (v Value) Uint16() uint16 {
	p, _ := ReadUint16(v)
	return p
}

func (v Value) Uint32() uint32 {
	p, _ := ReadUint32(v)
	return p
}

func (v Value) Uint64() uint64 {
	p, _ := ReadUint64(v)
	return p
}

func (v Value) Float32() float32 {
	p, _ := ReadFloat32(v)
	return p
}

func (v Value) Float64() float64 {
	p, _ := ReadFloat64(v)
	return p
}

func (v Value) Bytes() []byte {
	p, _ := ReadBytes(v)
	return p
}

func (v Value) String() string {
	p, _ := ReadString(v)
	return p
}

func (v Value) List() List {
	p, _ := GetList(v)
	return p
}

func (v Value) Message() Message {
	p, _ := GetMessage(v)
	return p
}
