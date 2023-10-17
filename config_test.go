package gocli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Field1 string
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

func TestGetConfigValue(t *testing.T) {
	t.Parallel()
	cfg := &TestConfig{Field1: "test", Field2: 42, Field3: true}
	cmd := NewCommand("test").WithConfig(cfg)
	cmdConfig := cmd.Config
	stringValue := GetConfigValue[string](cmdConfig, "field1")
	assert.Equal(t, "test", stringValue)
	intValue := GetConfigValue[int](cmdConfig, "field2")
	assert.Equal(t, 42, intValue)
	boolValue := GetConfigValue[bool](cmdConfig, "field3")
	assert.True(t, boolValue)
}

func TestCastConfig(t *testing.T) {
	t.Parallel()
	cfg := &TestConfig{Field1: "test", Field2: 42, Field3: true}
	cmd := NewCommand("test").WithConfig(cfg)
	cmdConfig := cmd.Config
	actual := CastConfig[TestConfig](cmdConfig)
	assert.True(t, reflect.DeepEqual(cfg, actual))
}

func BenchmarkGetConfigValue(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	config := &TestConfig{
		Field1: "test",
	}
	for i := 0; i < b.N; i++ {
		_ = GetConfigValue[string](config, "field1")
	}
}
