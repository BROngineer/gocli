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

func TestTypedConfig(t *testing.T) {
	t.Parallel()
	cfg := &TestConfig{Field1: "test", Field2: 42, Field3: true}
	cmd := NewCommand("test").WithConfig(cfg)
	cmdConfig := cmd.Config
	actual := TypedConfig[TestConfig](cmdConfig)
	assert.True(t, reflect.DeepEqual(cfg, actual))
}

func BenchmarkCastConfig(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	cfg := &TestConfig{Field1: "test", Field2: 42, Field3: true}
	cmd := NewCommand("test").WithConfig(cfg)
	raw := cmd.Config
	for i := 0; i < b.N; i++ {
		_ = TypedConfig[TestConfig](raw)
	}
}
