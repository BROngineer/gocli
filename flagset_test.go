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

func BenchmarkFlagSet_GetString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		set := NewFlagSet()
		set.AddFlag(NewFlag[string]("test", ""))
		flag := set.Flag("test")
		_ = flag.Parse("sample")
		_, _ = set.GetString("test")
	}
}

func BenchmarkFlagSet_GetInt(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		set := NewFlagSet()
		set.AddFlag(NewFlag[int]("test", ""))
		flag := set.Flag("test")
		_ = flag.Parse("42")
		_, _ = set.GetInt("test")
	}
}

func BenchmarkFlagSet_GetBool(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		set := NewFlagSet()
		set.AddFlag(NewFlag[bool]("test", ""))
		flag := set.Flag("test")
		_ = flag.Parse("true")
		_, _ = set.GetBool("test")
	}
}
