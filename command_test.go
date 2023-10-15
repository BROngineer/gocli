package gocli

import (
	"testing"
)

func BenchmarkNewCommand(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewCommand("sample")
	}
}

func BenchmarkCommand_WithSubcommand(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sub := NewCommand("subcommand")
		_ = NewCommand("sample").WithSubcommand(sub)
	}
}

func BenchmarkCommand_WithRunFunc(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := func(command Command) {}
		_ = NewCommand("sample").WithRunFunc(f)
	}
}

func BenchmarkCommand_WithRunEFunc(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := func(command Command) error { return nil }
		_ = NewCommand("sample").WithRunEFunc(f)
	}
}

func BenchmarkCommand_WithFlag(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := NewFlag[string]("test", "")
		_ = NewCommand("sample").WithFlag(f)
	}
}
