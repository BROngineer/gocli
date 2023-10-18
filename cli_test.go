package gocli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCLIApp(t *testing.T) {
	t.Parallel()
	cli := NewCLIApp()
	assert.NotNil(t, cli)
}

func TestCLIApp_Evaluate(t *testing.T) {
	var cmd Command
	var err error
	t.Parallel()
	cmd = NewCommand("test")
	cli := NewCLIApp()
	err = cli.Evaluate(cmd, []string{})
	assert.NoError(t, err)
	err = cli.Evaluate(cmd, []string{"-f"})
	assert.Error(t, err)
	cmd = NewCommand("test").WithFlag(NewFlag[string]("flag", ""))
	err = cli.Evaluate(cmd, []string{"test"})
	assert.Error(t, err)
}

func TestCLIApp_Execute(t *testing.T) {
	t.Parallel()
	cli := NewCLIApp()
	err := cli.Execute()
	assert.Error(t, err)
	cmd := NewCommand("test").WithRunFunc(func(command Command) {})
	err = cli.Evaluate(cmd, []string{})
	assert.NoError(t, err)
	err = cli.Execute()
	assert.NoError(t, err)
}

func TestInheritFlags(t *testing.T) {
	t.Parallel()
	cmd2 := NewCommand("sample")
	cmd1 := NewCommand("test").
		WithFlag(NewFlag[bool]("verbose", "").SetShared()).
		WithSubcommand(cmd2)
	inheritFlags(&cmd1)
	inherited := cmd2.Flag("verbose")
	assert.NotNil(t, inherited)
}

func TestValidateFlags(t *testing.T) {
	var err error
	var fs FlagSet
	t.Parallel()
	fs = NewFlagSet()
	fs.AddFlag(NewFlag[string]("flag", "").SetRequired())
	err = validateFlags(fs)
	assert.Error(t, err)
	fs = NewFlagSet()
	f := NewFlag[string]("flag", "")
	f.SetParsed()
	fs.AddFlag(f)
	err = validateFlags(fs)
	assert.NoError(t, err)
	fs = NewFlagSet()
	fs.AddFlag(NewFlag[string]("flag", ""))
	err = validateFlags(fs)
	assert.Error(t, err)
	fs = NewFlagSet()
	fs.AddFlag(NewFlag[string]("flag", "").WithDefault("Val"))
	err = validateFlags(fs)
	assert.NoError(t, err)
}

func TestSplitEqualsChar(t *testing.T) {
	t.Parallel()
	input := "flag=Val"
	flag, value := splitEqualsChar(input)
	assert.Equal(t, "flag", flag)
	assert.Equal(t, "Val", value)
}

func TestEvaluate(t *testing.T) {
	var cmd Command
	var err error
	t.Parallel()
	cmd = NewCommand("test").
		WithSubcommand(NewCommand("run").
			WithFlag(NewFlag[int]("flag", "")).
			WithFlag(NewFlag[bool]("sample", "")))
	_, err = evaluate(cmd, []string{"test", "run", "-flag=42", "--sample"})
	assert.NoError(t, err)
	_, err = evaluate(cmd, []string{"test", "run", "-flag", "42", "--sample"})
	assert.NoError(t, err)
	_, err = evaluate(cmd, []string{"test", "server"})
	assert.Error(t, err)
	_, err = evaluate(cmd, []string{"test", "run", "-flag", "-sample"})
	assert.Error(t, err)
	_, err = evaluate(cmd, []string{"test", "run", "-sample", "-flag"})
	assert.Error(t, err)
	_, err = evaluate(cmd, []string{"test", "run", "-sample", "-flag", "one"})
	assert.Error(t, err)
}
