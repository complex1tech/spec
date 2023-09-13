package generator

import "github.com/basecomplextech/spec/internal/lang/model"

func (w *writer) service(def *model.Definition) error {
	if err := w.iface(def); err != nil {
		return err
	}
	if err := w.client(def); err != nil {
		return err
	}
	// if err := w.handler(def); err != nil {
	// 	return err
	// }
	return nil
}

func (w *writer) iface(def *model.Definition) error {
	w.linef(`// %v`, def.Name)
	w.line()
	w.linef(`type %v interface {`, def.Name)

	for _, m := range def.Service.Methods {
		if err := w.ifaceMethod(def, m); err != nil {
			return err
		}
	}

	w.linef(`}`)
	w.line()
	return nil
}

func (w *writer) ifaceMethod(def *model.Definition, m *model.Method) error {
	if err := w.ifaceMethod_args(def, m); err != nil {
		return err
	}
	if err := w.ifaceMethod_results(def, m); err != nil {
		return err
	}
	w.line()
	return nil
}

func (w *writer) ifaceMethod_args(def *model.Definition, m *model.Method) error {
	methodName := toUpperCamelCase(m.Name)

	w.writef(`%v`, methodName)
	w.write(`(cancel <-chan struct{}, `)

	multi := false

	switch {
	case m.Input != nil:
		typeName := typeName(m.Input)
		w.writef(`req_ %v`, typeName)

	case m.InputFields != nil:
		fields := m.InputFields.List

		multi = len(fields) > 3
		if multi {
			w.line()
		}

		for _, field := range fields {
			argName := toLowerCameCase(field.Name)
			typeName := typeRefName(field.Type)

			if multi {
				w.linef(`%v_ %v, `, argName, typeName)
			} else {
				w.writef(`%v_ %v, `, argName, typeName)
			}
		}
	}

	w.write(`) `)
	return nil
}

func (w *writer) ifaceMethod_results(def *model.Definition, m *model.Method) error {
	w.write(`(`)

	out := m.Output
	multi := false

	switch {
	case m.Sub:
		typeName := typeName(out)
		w.writef(`_sub %v, `, typeName)

	case m.Output != nil:
		typeName := typeName(out)
		w.writef(`_resp %v, `, typeName)

	case m.OutputFields != nil:
		fields := m.OutputFields.List

		multi = len(fields) > 1
		if multi {
			w.line()
		}

		for _, field := range fields {
			resName := toLowerCameCase(field.Name)
			typeName := typeName(field.Type)

			if multi {
				w.linef(`_%v %v, `, resName, typeName)
			} else {
				w.writef(`_%v %v, `, resName, typeName)
			}
		}
	}

	if multi {
		w.line(`_st status.Status,`)
	} else {
		w.write(`_st status.Status`)
	}
	w.line(`)`)
	return nil
}
