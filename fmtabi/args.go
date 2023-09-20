package fmtabi

import (
	"fmt"
	"reflect"
	"strings"
)

func FormatArgs(args interface{}) string {
	return formatInterface(args, 1)
}

func formatInterface(value interface{}, callDepth int) string {
	switch value := value.(type) {
	case []byte:
		return formatBytes(value)
	case byte:
		return fmt.Sprintf("%#x", value)
	default:
		switch kind := reflect.TypeOf(value).Kind(); kind {
		case reflect.Struct:
			return formatStruct(value, callDepth)
		case reflect.Slice:
			return formatSlice(value, callDepth)
		default:
			return fmt.Sprintf("%v", value)
		}
	}
}

func formatBytes(bytes []byte) string {
	if len(bytes) == 0 {
		return "0x"
	}
	if len(bytes) < 33 {
		return fmt.Sprintf("%#x", bytes)
	}

	head := bytes[:16]
	tail := bytes[len(bytes)-16:]

	return fmt.Sprintf("%#x...%x", head, tail)
}

func formatStruct(obj interface{}, callDepth int) string {
	num := reflect.TypeOf(obj).NumField()
	fields := make([]string, 0, num)

	for i := 0; i < num; i++ {
		name := reflect.TypeOf(obj).Field(i).Tag.Get("json")
		value := formatInterface(reflect.ValueOf(obj).Field(i).Interface(), callDepth+1)
		fields = append(fields, name+": "+value)
	}
	content := strings.Join(fields, createSeparator(callDepth))
	return fmt.Sprintf("{\n%v%v\n%v}", createPadding(callDepth), content, createPadding(callDepth-1))
}

func formatSlice(arr interface{}, callDepth int) string {
	length := reflect.ValueOf(arr).Len()
	elems := make([]string, 0, length)

	for i := 0; i < length; i++ {
		elem := formatInterface(reflect.ValueOf(arr).Index(i).Interface(), callDepth+1)
		elems = append(elems, elem)
	}
	content := strings.Join(elems, createSeparator(callDepth))
	return fmt.Sprintf("[\n%v%v\n]", createPadding(callDepth), content)
}

func createSeparator(callDepth int) string {
	return ",\n" + createPadding(callDepth)
}

func createPadding(n int) string {
	if n < 1 {
		return ""
	}
	return strings.Repeat("  ", n)
}
