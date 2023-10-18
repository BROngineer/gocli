package gocli

import (
	"strconv"
	"strings"
	"time"
)

type allowed interface {
	~string | ~int | ~bool | ~[]string | time.Duration
}

type FlagValue interface {
	IsNil() bool
	Value() any
}

type Value[T allowed] struct {
	defined bool
	val     *T
}

func (v Value[T]) IsNil() bool {
	return !v.defined
}

func (v Value[T]) Value() any {
	return v.val
}

type GenericFlag[T allowed] struct {
	name        string
	shorthand   string
	description string
	DefVal      Value[T]
	Val         Value[T]
	shared      bool
	required    bool
	parsed      bool
}

func NewFlag[T allowed](name, description string) *GenericFlag[T] {
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
	f.DefVal = Value[T]{true, &value}
	return f
}

func (f *GenericFlag[T]) Name() string {
	return f.name
}

func (f *GenericFlag[T]) Shorthand() string {
	return f.shorthand
}

func (f *GenericFlag[T]) Value() FlagValue {
	return f.Val
}

func (f *GenericFlag[T]) ValueOrDefault() FlagValue {
	if !f.Val.IsNil() {
		return f.Val
	}
	return f.DefVal
}

func (f *GenericFlag[T]) Parse(in string) error {
	var v T
	switch f.Val.Value().(type) {
	case *string:
		v = any(in).(T)
	case *int:
		d, err := strconv.Atoi(in)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *bool:
		d, err := strconv.ParseBool(in)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *[]string:
		d := strings.Split(in, ",")
		v = any(d).(T)
	case *time.Duration:
		d, err := time.ParseDuration(in)
		if err != nil {
			return err
		}
		v = any(d).(T)
	}
	f.Val = Value[T]{true, &v}
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
