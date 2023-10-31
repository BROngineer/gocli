package gocli

import (
	"time"
)

func StringFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[string]{name: name}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func IntFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[int]{name: name}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func FloatFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[float64]{name: name}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func BoolFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[bool]{name: name, defVal: initValue[bool]()}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func SliceFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[[]string]{name: name}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func DurationFlag(name string, opts ...FlagOption) Flag {
	f := &genericFlag[time.Duration]{name: name}
	for _, opt := range opts {
		opt(f)
	}
	return f
}
