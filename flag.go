package gocli

import (
	"strconv"
)

type Flag interface {
	Name() string
	Value() any
	Parse(string) error
	Shared() bool
	SetShared()
	Required() bool
	SetRequired()
}

type FlagSet struct {
	Flags map[string]Flag
}

func NewFlagSet() FlagSet {
	return FlagSet{
		Flags: make(map[string]Flag),
	}
}

func (f *FlagSet) Merge(in *FlagSet) {
	for _, v := range in.Flags {
		f.AddFlag(v)
	}
}

func (f *FlagSet) AddFlag(flag Flag) {
	f.Flags[flag.Name()] = flag
}

type GenericFlag[T any] struct {
	name         string
	description  string
	defaultValue *T
	value        *T
	shared       bool
	required     bool
}

func NewFlag[T any](name, description string, defaultValue *T) *GenericFlag[T] {
	return &GenericFlag[T]{
		name:         name,
		description:  description,
		defaultValue: defaultValue}
}

func NewBooleanFlag[T any](name, description string) *GenericFlag[T] {
	defaultValue := any(false).(T)
	return &GenericFlag[T]{
		name:         name,
		description:  description,
		defaultValue: &defaultValue,
	}
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
		v := any(true).(T)
		f.value = &v
	}
	return nil
}

func (f *GenericFlag[T]) Shared() bool {
	return f.shared
}

func (f *GenericFlag[T]) SetShared() {
	f.shared = true
}

func (f *GenericFlag[T]) Required() bool {
	return f.required
}

func (f *GenericFlag[T]) SetRequired() {
	f.required = true
}
