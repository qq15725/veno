package refl

import "reflect"

func IndirectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func FieldTagValue(t reflect.Type, name string, tag string) string {
	if field, ok := t.FieldByName(name); ok {
		return field.Tag.Get(tag)
	}
	return ""
}
