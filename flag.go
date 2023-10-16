package gocli

import (
	"strconv"
)

type GenericFlag[T any] struct {
	name         string
	shorthand    string
	description  string
	defaultValue *T
	value        *T
	shared       bool
	required     bool
	parsed       bool
}

func NewFlag[T any](name, description string) *GenericFlag[T] {
	return &GenericFlag[T]{
		name:        name,
		description: description,
	}
}

func (f *GenericFlag[T]) WithShorthand(s string) *GenericFlag[T] {
	f.shorthand = s
	return f
}

func (f *GenericFlag[T]) WithDefault(value T) *GenericFlag[T] {
	f.defaultValue = &value
	return f
}

func (f *GenericFlag[T]) Name() string {
	return f.name
}

func (f *GenericFlag[T]) Value() any {
	if f.value == nil {
		return f.defaultValue
	}
	return f.value
}

func (f *GenericFlag[T]) Parse(in string) error {
	switch any(f.value).(type) {
	case *string:
		v := any(in).(T)
		f.value = &v
	case *int:
		d, err := strconv.Atoi(in)
		if err != nil {
			return err
		}
		v := any(d).(T)
		f.value = &v
	case *bool:
		d, err := strconv.ParseBool(in)
		if err != nil {
			return err
		}
		v := any(d).(T)
		f.value = &v
	}
	return nil
}

func (f *GenericFlag[T]) Shared() bool {
	return f.shared
}

func (f *GenericFlag[T]) SetShared() Flag {
	f.shared = true
	return f
}

func (f *GenericFlag[T]) Required() bool {
	return f.required
}

func (f *GenericFlag[T]) SetRequired() Flag {
	f.required = true
	return f
}

func (f *GenericFlag[T]) Parsed() bool {
	return f.parsed
}

func (f *GenericFlag[T]) SetParsed() {
	f.parsed = true
}
