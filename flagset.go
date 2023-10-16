package gocli

import (
	"fmt"
	"time"
)

type Flag interface {
	Name() string
	Shorthand() string
	Value() any
	Parse(string) error
	Shared() bool
	SetShared() Flag
	Required() bool
	SetRequired() Flag
	Parsed() bool
	SetParsed()
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

func (f *FlagSet) Flag(name string) Flag {
	flag, found := f.Flags[name]
	if found {
		return flag
	}
	for _, flag = range f.Flags {
		if flag.Shorthand() == name {
			return flag
		}
	}
	return nil
}

func (f *FlagSet) GetString(name string) (*string, error) {
	flag := f.Flag(name)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*string)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not a string type")
}

func (f *FlagSet) GetInt(name string) (*int, error) {
	flag := f.Flag(name)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*int)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not an int type")
}

func (f *FlagSet) GetBool(name string) (*bool, error) {
	flag := f.Flag(name)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*bool)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not a boolean type")
}

func (f *FlagSet) GetStringSlice(name string) (*[]string, error) {
	flag := f.Flag(name)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*[]string)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not a string slice type")
}

func (f *FlagSet) GetDuration(name string) (*time.Duration, error) {
	flag := f.Flag(name)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*time.Duration)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not a time.Duration type")
}
