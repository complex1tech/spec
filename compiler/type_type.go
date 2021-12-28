package compiler

import (
	"fmt"

	"github.com/baseone-run/spec/parser"
)

type Kind string

const (
	// builtin

	KindUndefined Kind = ""
	KindBool      Kind = "bool"

	KindInt8  Kind = "int8"
	KindInt16 Kind = "int16"
	KindInt32 Kind = "int32"
	KindInt64 Kind = "int64"

	KindUint8  Kind = "uint8"
	KindUint16 Kind = "uint16"
	KindUint32 Kind = "uint32"
	KindUint64 Kind = "uint64"

	KindFloat32 Kind = "float32"
	KindFloat64 Kind = "float64"

	KindBytes  Kind = "bytes"
	KindString Kind = "string"

	// references

	KindReference Kind = "reference"
	KindImport    Kind = "import"
	KindList      Kind = "list"
	KindNullable  Kind = "nullable"
)

var builtin = map[string]*Type{
	string(KindBool): newBuiltinType(KindBool),

	string(KindInt8):  newBuiltinType(KindInt8),
	string(KindInt16): newBuiltinType(KindInt16),
	string(KindInt32): newBuiltinType(KindInt32),
	string(KindInt64): newBuiltinType(KindInt64),

	string(KindUint8):  newBuiltinType(KindUint8),
	string(KindUint16): newBuiltinType(KindUint16),
	string(KindUint32): newBuiltinType(KindUint32),
	string(KindUint64): newBuiltinType(KindUint64),

	string(KindFloat32): newBuiltinType(KindFloat32),
	string(KindFloat64): newBuiltinType(KindFloat64),

	string(KindBytes):  newBuiltinType(KindBytes),
	string(KindString): newBuiltinType(KindString),
}

type Type struct {
	Kind       Kind
	Name       string
	Element    *Type  // element type in list, reference and nullable types
	ImportName string // imported package name, "pkg" in "pkg.Type"

	// Resolved
	Ref    *Definition
	Import *Import

	// Utility flags
	Builtin    bool
	Imported   bool
	List       bool
	Nullable   bool
	Referenced bool
}

func newType(ptype *parser.Type) (*Type, error) {
	switch ptype.Kind {
	case parser.KindBase:
		type_, ok := builtin[ptype.Name]
		if ok {
			return type_, nil
		}

		type_ = &Type{
			Kind: KindReference,
			Name: ptype.Name,

			Referenced: true,
		}
		return type_, nil

	case parser.KindImport:
		type_ := &Type{
			Kind:       KindImport,
			Name:       ptype.Name,
			ImportName: ptype.Import,

			Imported: true,
		}
		return type_, nil

	case parser.KindNullable:
		elem, err := newType(ptype.Element)
		if err != nil {
			return nil, err
		}

		type_ := &Type{
			Kind:    KindNullable,
			Name:    "*",
			Element: elem,

			Nullable: true,
		}
		return type_, nil

	case parser.KindList:
		elem, err := newType(ptype.Element)
		if err != nil {
			return nil, err
		}

		type_ := &Type{
			Kind:    KindList,
			Name:    "[]",
			Element: elem,

			List: true,
		}
		return type_, nil
	}

	return nil, fmt.Errorf("unsupported type kind, kind=%v, name=%v", ptype.Kind, ptype.Name)
}

func newBuiltinType(kind Kind) *Type {
	return &Type{
		Kind:    kind,
		Name:    string(kind),
		Builtin: true,
	}
}