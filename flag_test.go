package gocli

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFlag(t *testing.T) {
	t.Parallel()
	f := NewFlag[string]("test", "")
	assert.NotNil(t, f)
}

func TestGenericFlag_Name(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "").Name()
		assert.Equal(t, names[i], n)
	}
}

func TestGenericFlag_WithDefault(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "").WithDefault(names[i])
		v := n.ValueOrDefault().(*string)
		assert.NotNil(t, v)
		assert.Equal(t, names[i], *v)
	}
}

func TestGenericFlag_Shared(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "").SetShared()
		v := n.Shared()
		assert.True(t, v)
	}
}

func TestGenericFlag_Required(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "").SetRequired()
		v := n.Required()
		assert.True(t, v)
	}
}

func TestGenericFlag_Parsed(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "")
		n.SetParsed()
		v := n.Parsed()
		assert.True(t, v)
	}
}

func TestGenericFlag_Shorthand(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "").WithShorthand(string(names[i][0]))
		v := n.Shorthand()
		assert.Equal(t, string(names[i][0]), v)
	}
}

func TestGenericFlag_ParseString(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[string](names[i], "")
		err := n.Parse(names[i])
		assert.NoError(t, err)
		v1 := n.Value().(*string)
		v2 := n.ValueOrDefault().(*string)
		assert.Equal(t, names[i], *v1)
		assert.Equal(t, names[i], *v2)
		assert.Equal(t, *v1, *v2)
	}
}

func TestGenericFlag_ParseInt(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[int](names[i], "")
		err := n.Parse("10s")
		assert.Error(t, err)
		err = n.Parse(strconv.Itoa(i))
		assert.NoError(t, err)
		v1 := n.Value().(*int)
		v2 := n.ValueOrDefault().(*int)
		assert.Equal(t, i, *v1)
		assert.Equal(t, i, *v2)
		assert.Equal(t, *v1, *v2)
	}
}

func TestGenericFlag_ParseBool(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[bool](names[i], "")
		err := n.Parse(names[i])
		assert.Error(t, err)
		err = n.Parse("true")
		assert.NoError(t, err)
		v1 := n.Value().(*bool)
		v2 := n.ValueOrDefault().(*bool)
		assert.True(t, *v1)
		assert.True(t, *v2)
		assert.Equal(t, *v1, *v2)
	}
}

func TestGenericFlag_ParseStringSlice(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		value := strings.Join(names, ",")
		n := NewFlag[[]string](names[i], "")
		err := n.Parse(value)
		assert.NoError(t, err)
		v1 := n.Value().(*[]string)
		v2 := n.ValueOrDefault().(*[]string)
		assert.True(t, reflect.DeepEqual(names, *v1))
		assert.True(t, reflect.DeepEqual(names, *v2))
		assert.True(t, reflect.DeepEqual(*v1, *v2))
	}
}

func TestGenericFlag_ParseDuration(t *testing.T) {
	t.Parallel()
	names := []string{"test", "flag", "sample"}
	for i := 0; i < len(names); i++ {
		n := NewFlag[time.Duration](names[i], "")
		err := n.Parse("20")
		assert.Error(t, err)
		err = n.Parse("10s")
		assert.NoError(t, err)
		v1 := n.Value().(*time.Duration)
		v2 := n.ValueOrDefault().(*time.Duration)
		assert.Equal(t, time.Second*10, *v1)
		assert.Equal(t, time.Second*10, *v2)
		assert.Equal(t, *v1, *v2)
	}
}

func BenchmarkNewFlag(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlag[string]("test", "")
	}
}

func BenchmarkGenericFlag_WithDefault(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlag[string]("test", "").WithDefault("empty")
	}
}

func BenchmarkGenericFlag_SetShared(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlag[string]("test", "").SetShared()
	}
}

func BenchmarkGenericFlag_SetRequired(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlag[string]("test", "").SetRequired()
	}
}

func BenchmarkGenericFlag_SetParsed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewFlag[string]("test", "").SetParsed()
	}
}

func BenchmarkGenericFlag_Parse(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	f := NewFlag[string]("test", "")
	for i := 0; i < b.N; i++ {
		_ = f.Parse("test")
	}
}
