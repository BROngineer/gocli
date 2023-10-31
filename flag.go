package gocli

import (
	"strconv"
	"strings"
	"time"
)

type allowed interface {
	~string | ~int | ~float64 | ~bool | ~[]string | time.Duration
}

type flagValue interface {
	isNil() bool
	unwrap() any
}

type Value[T allowed] struct {
	defined bool
	val     *T
}

func initValue[T allowed]() Value[T] {
	return Value[T]{
		defined: true,
		val:     new(T),
	}
}

func (v Value[T]) isNil() bool {
	return !v.defined
}

func (v Value[T]) unwrap() any {
	return v.val
}

type Flag interface {
	flagDataGetter
	flagDataSetter
}

type flagDataGetter interface {
	Name() string
	Description() string
	Shorthand() string
	Required() bool
	Shared() bool
	Parsed() bool
	Value() any
	IsNilValue() bool
}

type flagDataSetter interface {
	setDescription(string)
	setShorthand(string)
	setRequired()
	setShared()
	setDefVal(flagValue)
	parse(string) error
}

type genericFlag[T allowed] struct {
	name        string
	description string
	shorthand   string
	required    bool
	shared      bool
	parsed      bool
	val         Value[T]
	defVal      Value[T]
}

func (f *genericFlag[T]) Name() string {
	return f.name
}

func (f *genericFlag[T]) Description() string {
	return f.description
}

func (f *genericFlag[T]) setDescription(value string) {
	f.description = value
}

func (f *genericFlag[T]) Shorthand() string {
	return f.shorthand
}

func (f *genericFlag[T]) setShorthand(value string) {
	f.shorthand = value
}

func (f *genericFlag[T]) Required() bool {
	return f.required
}

func (f *genericFlag[T]) setRequired() {
	f.required = true
}

func (f *genericFlag[T]) Shared() bool {
	return f.shared
}

func (f *genericFlag[T]) setShared() {
	f.shared = true
}

func (f *genericFlag[T]) Parsed() bool {
	return f.parsed
}

func (f *genericFlag[T]) parse(input string) error {
	var v T
	switch f.val.unwrap().(type) {
	case *string:
		v = any(input).(T)
	case *int:
		d, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *float64:
		d, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *bool:
		d, err := strconv.ParseBool(input)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *[]string:
		d := strings.Split(input, ",")
		v = any(d).(T)
	case *time.Duration:
		d, err := time.ParseDuration(input)
		if err != nil {
			return err
		}
		v = any(d).(T)
	}
	f.val = Value[T]{defined: true, val: &v}
	f.parsed = true
	return nil
}

func (f *genericFlag[T]) value() flagValue {
	if f.val.defined {
		return f.val
	}
	return f.defVal
}

func (f *genericFlag[T]) Value() any {
	return f.value().unwrap()
}

func (f *genericFlag[T]) setDefVal(value flagValue) {
	f.defVal = value.(Value[T])
}

func (f *genericFlag[T]) IsNilValue() bool {
	return f.value().isNil()
}
