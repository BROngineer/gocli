package gocli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommandConfig struct {
	Field1 string
	Field2 int
	Field3 bool
}

func (t *CommandConfig) LoadFromFiles(_ []string) error {
	return nil
}

func (t *CommandConfig) LoadFromEnv(_ string) error {
	return nil
}

func (t *CommandConfig) LoadFromFlags(_ FlagSet) error {
	return nil
}

func TestNewCommand(t *testing.T) {
	t.Parallel()
	cmd := NewCommand("test")
	assert.NotEmpty(t, cmd)
	assert.Equal(t, "test", cmd.Name)
	assert.NotEmpty(t, cmd.FlagSet)
}

func TestCommand_WithFlag(t *testing.T) {
	t.Parallel()
	flag := NewFlag[string]("test", "")
	cmd := NewCommand("test").WithFlag(flag)
	actual := cmd.Flag("test")
	assert.NotNil(t, actual)
	assert.True(t, reflect.DeepEqual(flag, actual))
	assert.Equal(t, len(cmd.Flags()), 1)
}

func TestCommand_WithSubcommand(t *testing.T) {
	t.Parallel()
	cmd := NewCommand("test")
	assert.Nil(t, cmd.Subcommands)
	subcmd := NewCommand("subtest")
	cmd = cmd.WithSubcommand(subcmd)
	assert.NotNil(t, cmd.Subcommands)
}

func TestCommand_WithRunFunc(t *testing.T) {
	t.Parallel()
	cmd := NewCommand("test").
		WithRunFunc(func(command Command) {

		})
	assert.NotNil(t, cmd.Run)
}

func TestCommand_WithRunEFunc(t *testing.T) {
	t.Parallel()
	cmd := NewCommand("test").
		WithRunEFunc(func(command Command) error {
			return nil
		})
	assert.NotNil(t, cmd.RunE)
}

func TestCommand_WithConfig(t *testing.T) {
	t.Parallel()
	cfg := &CommandConfig{}
	cmd := NewCommand("test").
		WithConfig(cfg)
	assert.NotNil(t, cmd.Config)
}

func TestCommand_Execute(t *testing.T) {
	var err error
	t.Parallel()
	cmd := NewCommand("test")
	err = cmd.Execute()
	assert.Error(t, err)
	cmd = cmd.WithRunFunc(func(command Command) {})
	err = cmd.Execute()
	assert.NoError(t, err)
	cmd = cmd.WithRunEFunc(func(command Command) error { return nil })
	err = cmd.Execute()
	assert.NoError(t, err)
}

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
