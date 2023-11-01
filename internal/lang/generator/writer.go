package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/basecomplextech/spec/internal/lang/model"
)

type writer struct {
	b bytes.Buffer

	skipRPC bool
}

func newWriter(skipRPC bool) *writer {
	return &writer{
		b: bytes.Buffer{},

		skipRPC: skipRPC,
	}
}

func (w *writer) line(args ...string) {
	w.write(args...)
	w.b.WriteString("\n")
}

func (w *writer) linef(format string, args ...interface{}) {
	w.writef(format, args...)
	w.b.WriteString("\n")
}

func (w *writer) write(args ...string) {
	for _, s := range args {
		w.b.WriteString(s)
	}
}

func (w *writer) writef(format string, args ...interface{}) {
	if len(args) == 0 {
		w.write(format)
		return
	}

	s := fmt.Sprintf(format, args...)
	w.b.WriteString(s)
}

func (w *writer) file(file *model.File) error {
	return newFileWriter(w).file(file)
}

// internal

func importPackage(imp *model.Import) string {
	pkg, ok := imp.Package.OptionNames[OptionPackage]
	if ok {
		return pkg.Value
	}

	return imp.ID
}

func toUpperCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		part = strings.ToLower(part)
		part = strings.Title(part)
		parts[i] = part
	}
	return strings.Join(parts, "")
}

func toLowerCameCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = toUpperCamelCase(s)
	return strings.ToLower(s[:1]) + s[1:]
}
