package gocli

import (
	"testing"
)

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
