package gocli

import (
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type FlagComparisonAssertion func(*testing.T, Flag, Flag)
type FlagValueAssertion func(*testing.T, any, any)

func CompareFlags(t *testing.T, expect, actual Flag) {
	assert.Equal(t, expect.Name(), actual.Name())
	assert.Equal(t, expect.Description(), actual.Description())
	assert.Equal(t, expect.Shorthand(), actual.Shorthand())
	assert.Equal(t, expect.Required(), actual.Required())
	assert.Equal(t, expect.Shared(), actual.Shared())
	assert.Equal(t,
		reflect.TypeOf(expect.Value()),
		reflect.TypeOf(actual.Value()))
}

func AssertFlagValue[T allowed](t *testing.T, expect, actual any) {
	actualTyped := actual.(*T)
	assert.Equal(t, reflect.TypeOf(expect), reflect.TypeOf(*actualTyped))
	assert.Equal(t, expect, *actualTyped)
}

func TestNewFlag(t *testing.T) {
	var defVal = "default"
	t.Parallel()
	tests := []struct {
		name      string
		expected  Flag
		actual    Flag
		assertion FlagComparisonAssertion
	}{
		{
			"string flag",
			&genericFlag[string]{name: "sample", description: "sample"},
			StringFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"int flag",
			&genericFlag[int]{name: "sample", description: "sample"},
			IntFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"float flag",
			&genericFlag[float64]{name: "sample", description: "sample"},
			FloatFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"bool flag",
			&genericFlag[bool]{name: "sample", description: "sample"},
			BoolFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"slice flag",
			&genericFlag[[]string]{name: "sample", description: "sample"},
			SliceFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"duration flag",
			&genericFlag[time.Duration]{name: "sample", description: "sample"},
			DurationFlag("sample", FlagDescription("sample")),
			CompareFlags,
		}, {
			"with options",
			&genericFlag[string]{
				name:        "sample",
				description: "sample",
				shorthand:   "s",
				required:    true,
				shared:      true,
				defVal: Value[string]{
					defined: true,
					val:     &defVal,
				},
			},
			StringFlag("sample",
				FlagDescription("sample"),
				Shorthand("s"),
				Required(),
				Shared(),
				Default[string]("default")),
			CompareFlags,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expected, tt.actual)
		})
	}
}

func TestParseFlag(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		value     string
		expected  any
		flag      Flag
		assertion FlagValueAssertion
	}{
		{
			"parse string",
			"test",
			"test",
			StringFlag("sample"),
			AssertFlagValue[string],
		}, {
			"parse int",
			"42",
			42,
			IntFlag("sample"),
			AssertFlagValue[int],
		}, {
			"parse float",
			"1.0",
			1.0,
			FloatFlag("sample"),
			AssertFlagValue[float64],
		}, {
			"parse bool",
			"true",
			true,
			BoolFlag("sample"),
			AssertFlagValue[bool],
		}, {
			"parse slice",
			"a,b,c",
			[]string{"a", "b", "c"},
			SliceFlag("sample"),
			AssertFlagValue[[]string],
		}, {
			"parse duration",
			"10s",
			time.Second * 10,
			DurationFlag("sample"),
			AssertFlagValue[time.Duration],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flag.parse(tt.value)
			assert.NoError(t, err)
			assert.True(t, true, flag.Parsed())
			actual := tt.flag.Value()
			tt.assertion(t, tt.expected, actual)
		})
	}
}

func TestParseFlagErr(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value string
		flag  Flag
	}{
		{
			"parse int",
			"one",
			IntFlag("sample"),
		}, {
			"parse float",
			"",
			FloatFlag("sample"),
		}, {
			"parse bool",
			"10",
			BoolFlag("sample"),
		}, {
			"parse duration",
			"10",
			DurationFlag("sample"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flag.parse(tt.value)
			assert.Error(t, err)
		})
	}
}
