package container

import (
	"errors"
	"reflect"
)

type Container struct {
	instances map[string]interface{}
	bindings  map[string]interface{}
}

func (ctr *Container) Bind(construct interface{}) error {
	typeOf := reflect.TypeOf(construct)

	if typeOf == nil {
		return errors.New("can't bind an untyped nil")
	}

	if typeOf.Kind() != reflect.Func {
		return errors.New("xassx")
	}

	for i := 0; i < typeOf.NumOut(); i++ {
		typeOfOut := typeOf.Out(i)

		abstract := typeOfOut.String()

		ctr.bindings[abstract] = construct
	}

	return nil
}

func (ctr *Container) Get(abstract interface{}) interface{} {
	return ctr.resolve(abstract)
}

func (ctr *Container) resolve(abstract interface{}) interface{} {
	typeOf := reflect.TypeOf(abstract)

	if typeOf == nil {
		return errors.New("can't resolve an untyped nil")
	}

	name := typeOf.String()

	if concrete, ok := ctr.instances[name]; ok {
		return concrete
	}

	concrete, ok := ctr.bindings[name]

	if !ok {
		return nil
	}

	valueOf := reflect.ValueOf(concrete)
	args := make([]reflect.Value, 0)
	values := valueOf.Call(args)
	for i := 0; i < len(values); i++ {
		value := values[i]
		ctr.instances[value.Type().String()] = value.Interface()
	}

	return ctr.instances[name]
}

func New() *Container {
	return &Container{
		instances: make(map[string]interface{}),
		bindings:  make(map[string]interface{}),
	}
}
