package gocli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlagSet(t *testing.T) {
	t.Parallel()
	fs := NewFlagSet()
	assert.NotEmpty(t, fs)
	assert.NotNil(t, fs.Flags)
}

func TestFlagSet_AddFlag(t *testing.T) {
	t.Parallel()
	fs := NewFlagSet()
	expected := NewFlag[string]("test", "")
	fs.AddFlag(NewFlag[string]("test", ""))
	actual := fs.Flag("test")
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestFlagSet_Merge(t *testing.T) {
	t.Parallel()
	fs1 := NewFlagSet()
	fs2 := NewFlagSet()
	fs1.AddFlag(NewFlag[string]("test", ""))
	fs2.Merge(&fs1)
	expected := fs1.Flag("test")
	actual := fs2.Flag("test")
	assert.NotNil(t, actual)
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestFlagSet_Flag(t *testing.T) {
	var actual Flag
	t.Parallel()
	fs := NewFlagSet()
	f := NewFlag[string]("test", "").WithShorthand("t")
	fs.AddFlag(f)
	actual = fs.Flag("test")
	assert.NotNil(t, actual)
	assert.True(t, reflect.DeepEqual(f, actual))
	actual = fs.Flag("t")
	assert.NotNil(t, actual)
	assert.True(t, reflect.DeepEqual(f, actual))
	actual = fs.Flag("testing")
	assert.Nil(t, actual)
}

func TestGetValue(t *testing.T) {
	t.Parallel()
	fs := NewFlagSet()
	fs.AddFlag(NewFlag[string]("test", "").WithDefault("sample"))
	v, err := GetValue[string](fs, "test")
	assert.NoError(t, err)
	assert.Equal(t, "sample", *v)
	_, err = GetValue[string](fs, "testing")
	assert.Error(t, err)
	_, err = GetValue[int](fs, "test")
	assert.Error(t, err)
}

func BenchmarkNewFlagSet(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlagSet()
	}
}

func BenchmarkFlagSet_AddFlag(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		set := NewFlagSet()
		set.AddFlag(NewFlag[string]("test", ""))
	}
}

func BenchmarkFlagSet_Merge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	set1 := NewFlagSet()
	set1.AddFlag(NewFlag[string]("test", ""))
	for i := 0; i < b.N; i++ {
		set2 := NewFlagSet()
		set2.Merge(&set1)
	}
}

func BenchmarkFlagSet_Flag(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	set1 := NewFlagSet()
	set1.AddFlag(NewFlag[string]("test", ""))
	for i := 0; i < b.N; i++ {
		_ = set1.Flag("test")
	}
}

func BenchmarkGetValue(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	set := NewFlagSet()
	set.AddFlag(NewFlag[string]("test-string", ""))
	set.AddFlag(NewFlag[int]("test-int", ""))
	set.AddFlag(NewFlag[bool]("test-bool", ""))
	for i := 0; i < b.N; i++ {
		_, _ = GetValue[string](set, "test-string")
		_, _ = GetValue[int](set, "test-int")
		_, _ = GetValue[bool](set, "test-bool")
	}
}
