package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Encode(buf *bytes.Buffer, v interface{}) error {
	return encode(buf, reflect.ValueOf(v), 0)
}

func indent(buf *bytes.Buffer, indentLevel int) {
	for i := 0; i < (indentLevel << 2); i++ {
		buf.WriteByte(' ')
	}
}

func newLine(buf *bytes.Buffer) {
	buf.WriteByte('\n')
}

func encode(buf *bytes.Buffer, v reflect.Value, indentLevel int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem(), indentLevel)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				indent(buf, indentLevel)
			}
			if err := encode(buf, v.Index(i), indentLevel+1); err != nil {
				return err
			}
			if i < v.Len()-1 {
				newLine(buf)
			}
		}
		buf.WriteByte(')')
		newLine(buf)
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				indent(buf, indentLevel)
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indentLevel+1); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i < v.NumField()-1 {
				newLine(buf)
			}
		}
		buf.WriteByte(')')
		newLine(buf)
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				indent(buf, indentLevel)
			}
			buf.WriteByte('(')
			if err := encode(buf, key, indentLevel+1); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indentLevel+1); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i < v.Len()-1 {
				newLine(buf)
			}
		}
		buf.WriteByte(')')
		newLine(buf)
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%g %g)", real(v.Complex()), imag(v.Complex()))
	default:
		return fmt.Errorf("unsupported type %s: %s", v.Type().Name())
	}

	return nil
}
