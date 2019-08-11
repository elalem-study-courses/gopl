package display

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/template"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatComposite(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s.type = %s", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Int8:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func formatComposite(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		var buf bytes.Buffer
		templateString := `{{.Name}} {
	{{range $name, $value := .Fields}}
	{{$name}} = {{$value}}
	{{end}}
}`
		template := template.Must(template.New("template").Parse(templateString))
		structInfo := struct {
			Name   string
			Fields map[string]string
		}{Fields: make(map[string]string)}

		structInfo.Name = v.Type().Name()
		for i := 0; i < v.NumField(); i++ {
			structInfo.Fields[v.Type().Field(i).Name] = formatComposite(v.Field(i))
		}
		if err := template.Execute(&buf, structInfo); err != nil {
			return fmt.Sprintf("template parsing: %v", err)
		}
		return buf.String()
	case reflect.Array, reflect.Slice:
		str := v.Type().Name() + "{"
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				str += ","
			}
			str += formatComposite(v.Index(i))
		}
		str += "}"
		return str
	default:
		return formatAtom(v)
	}
}
