package refl

import "reflect"

func IndirectValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func SetFieldValue(v reflect.Value, name string, value interface{}) reflect.Value {
	field := v.FieldByName(name)
	if field.CanSet() {
		field.Set(reflect.ValueOf(value))
	}
	return v
}
