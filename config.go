package gocli

import (
	"reflect"
	"strings"
)

type Config interface {
	LoadFromFiles([]string)
	LoadFromEnv(string)
	LoadFromFlags(FlagSet)
}

func GetConfigValue[T any](cfg Config, field string) T {
	var v T
	r := reflect.ValueOf(cfg)
	t := reflect.Indirect(r).Type()
	for i := 0; i < reflect.Indirect(r).NumField(); i++ {
		n := t.Field(i).Name
		if strings.ToLower(n) == strings.ToLower(field) {
			f := reflect.Indirect(r).FieldByName(n)
			if f.CanInterface() {
				v = f.Interface().(T)
				return v
			}
		}
	}
	return v
}
