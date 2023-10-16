package gocli

import (
	"testing"
)

func BenchmarkNewFlag(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewFlag[string]("test", "")
	}
}

func BenchmarkGenericFlag_SetDefault(b *testing.B) {
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
	flags := []struct {
		f Flag
		v string
	}{
		{
			NewFlag[string]("test", ""),
			"test",
		}, {
			NewFlag[int]("test", ""),
			"1",
		}, {
			NewFlag[bool]("test", ""),
			"true",
		},
	}
	for i := 0; i < b.N; i++ {
		for _, item := range flags {
			_ = item.f.Parse(item.v)
		}
	}
}
