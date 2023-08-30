package generator

import (
	"fmt"
	"strings"

	"github.com/basecomplextech/spec/internal/lang/model"
)

func (w *writer) enum(def *model.Definition) error {
	if err := w.enumDef(def); err != nil {
		return err
	}
	if err := w.enumValues(def); err != nil {
		return err
	}
	if err := w.newEnum(def); err != nil {
		return err
	}
	if err := w.parseEnum(def); err != nil {
		return err
	}
	if err := w.writeEnum(def); err != nil {
		return err
	}
	if err := w.enumString(def); err != nil {
		return err
	}
	return nil
}

func (w *writer) enumDef(def *model.Definition) error {
	w.linef(`// %v`, def.Name)
	w.line()
	w.linef("type %v int32", def.Name)
	w.line()
	return nil
}

func (w *writer) enumValues(def *model.Definition) error {
	w.line("const (")

	for _, val := range def.Enum.Values {
		// EnumValue Enum = 1
		name := enumValueName(val)
		w.linef("%v %v = %d", name, def.Name, val.Number)
	}

	w.line(")")
	w.line()
	return nil
}

func (w *writer) newEnum(def *model.Definition) error {
	name := def.Name
	w.linef(`func New%v(b []byte) %v {`, name, name)
	w.linef(`v, _, _ := encoding.DecodeInt32(b)`)
	w.linef(`return %v(v)`, name)
	w.linef(`}`)
	w.line()
	return nil
}

func (w *writer) parseEnum(def *model.Definition) error {
	name := def.Name
	w.linef(`func Parse%v(b []byte) (result %v, size int, err error) {`, name, name)
	w.linef(`v, size, err := encoding.DecodeInt32(b)`)
	w.linef(`if err != nil || size == 0 {
		return
	}`)
	w.linef(`result = %v(v)`, name)
	w.line(`return`)
	w.linef(`}`)
	w.line()
	return nil
}

func (w *writer) writeEnum(def *model.Definition) error {
	w.linef(`func Write%v(b buffer.Buffer, v %v) (int, error) {`, def.Name, def.Name)
	w.linef(`return encoding.EncodeInt32(b, int32(v))`)
	w.linef(`}`)
	w.line()
	return nil
}

func (w *writer) enumString(def *model.Definition) error {
	w.linef("func (e %v) String() string {", def.Name)
	w.line("switch e {")

	for _, val := range def.Enum.Values {
		name := enumValueName(val)
		w.linef("case %v:", name)
		w.linef(`return "%v"`, strings.ToLower(val.Name))
	}

	w.line("}")
	w.line(`return ""`)
	w.line("}")
	w.line()
	return nil
}

func enumValueName(val *model.EnumValue) string {
	name := toUpperCamelCase(val.Name)
	return fmt.Sprintf("%v_%v", val.Enum.Def.Name, name)
}