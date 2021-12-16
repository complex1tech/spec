package spec

// Type specifies a value type.
type Type byte

const (
	TypeNil   Type = 00
	TypeTrue  Type = 01
	TypeFalse Type = 02

	TypeInt8  Type = 10
	TypeInt16 Type = 11
	TypeInt32 Type = 12
	TypeInt64 Type = 13

	TypeUInt8  Type = 20
	TypeUInt16 Type = 21
	TypeUInt32 Type = 22
	TypeUInt64 Type = 23

	TypeFloat32 Type = 30
	TypeFloat64 Type = 31

	TypeBytes  Type = 40
	TypeString Type = 41

	TypeList    Type = 50
	TypeMessage Type = 60
)
