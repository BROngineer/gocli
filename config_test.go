package gocli

import (
	"testing"
)

type TestConfig struct {
	field1 string
	Field2 int
	Field3 bool
}

func (t *TestConfig) LoadFromFiles(_ []string) error {
	return nil
}

func (t *TestConfig) LoadFromEnv(_ string) error {
	return nil
}

func (t *TestConfig) LoadFromFlags(_ FlagSet) error {
	return nil
}

func BenchmarkGetConfigValue(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	config := &TestConfig{
		field1: "test",
	}
	for i := 0; i < b.N; i++ {
		_ = GetConfigValue[string](config, "field1")
	}
}
