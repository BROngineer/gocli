package gocli

import (
	"fmt"
	"time"
)

type Flag interface {
	Name() string
	Shorthand() string
	Value() any
	ValueOrDefault() any
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

func GetValue[T ~string | int | bool | []string | time.Duration](flagSet FlagSet, flagName string) (*T, error) {
	flag := flagSet.Flag(flagName)
	if flag == nil {
		return nil, fmt.Errorf("no flag found")
	}
	v, ok := flag.ValueOrDefault().(*T)
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("not a %T type", flag.Value())
}
