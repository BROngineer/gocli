package gocli

import (
	"reflect"
	"strings"
)

type Config interface {
	LoadFromFiles(paths []string) error
	LoadFromEnv(prefix string) error
	LoadFromFlags(flags FlagSet) error
}

func GetConfigValue[T any](cfg Config, field string) T {
	var v T
	var r reflect.Value
	r = reflect.Indirect(reflect.ValueOf(cfg))
	for i := 0; i < r.NumField(); i++ {
		n := r.Type().Field(i).Name
		if strings.EqualFold(field, n) {
			if r.FieldByName(n).CanInterface() {
				v = r.FieldByName(n).Interface().(T)
				break
			}
		}
	}
	return v
}

func CastConfig[T any](cfg Config) *T {
	return any(cfg).(*T)
}
