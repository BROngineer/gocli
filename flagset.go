package gocli

import (
	"fmt"
)

type Flag interface {
	Name() string
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
	return f.Flags[name]
}

func (f *FlagSet) GetString(name string) (*string, error) {
	flag, found := f.Flags[name]
	if !found {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*string)
	if !ok {
		return nil, fmt.Errorf("not a string type")
	}
	return v, nil
}

func (f *FlagSet) GetInt(name string) (*int, error) {
	flag, found := f.Flags[name]
	if !found {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*int)
	if !ok {
		return nil, fmt.Errorf("not an int type")
	}
	return v, nil
}

func (f *FlagSet) GetBool(name string) (*bool, error) {
	flag, found := f.Flags[name]
	if !found {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.Value().(*bool)
	if !ok {
		return nil, fmt.Errorf("not a boolean type")
	}
	return v, nil
}
