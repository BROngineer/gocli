package cli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommandComparisonAssertion func(*testing.T, *Command, *Command)

func CompareCommands(t *testing.T, expected, actual *Command) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.True(t, reflect.DeepEqual(expected.Flags, actual.Flags))
	assert.True(t, reflect.DeepEqual(expected.Subcommands, actual.Subcommands))
}

func SampleRun(_ *Command)        {}
func SampleRunE(_ *Command) error { return nil }

func TestNewCommand(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		expected  *Command
		actual    *Command
		assertion CommandComparisonAssertion
	}{
		{
			"orphan command",
			&Command{Name: "sample", Flags: make(FlagsSet), Subcommands: make(CommandsSet), Description: "test command"},
			NewCommand("sample", CommandDescription("test command")),
			CompareCommands,
		}, {
			"command with flags",
			&Command{Name: "sample", Flags: FlagsSet{"flag": &genericFlag[string]{name: "flag"}}, Subcommands: make(CommandsSet)},
			NewCommand("sample", Flags(StringFlag("flag"))),
			CompareCommands,
		}, {
			"command with subcommands",
			&Command{
				Name:  "sample",
				Flags: make(FlagsSet),
				Subcommands: CommandsSet{
					"sub": {
						Name:        "sub",
						Flags:       make(FlagsSet),
						Subcommands: make(CommandsSet),
					},
				},
			},
			NewCommand("sample", Subcommands(NewCommand("sub"))),
			CompareCommands,
		}, {
			"complex command",
			&Command{
				Name:  "sample",
				Flags: FlagsSet{"parent": &genericFlag[bool]{name: "parent"}},
				Subcommands: CommandsSet{
					"sub": {
						Name:        "sub",
						Flags:       FlagsSet{"flag": &genericFlag[string]{name: "flag"}},
						Subcommands: make(CommandsSet),
					},
				},
			},
			NewCommand("sample",
				Flags(BoolFlag("parent")),
				Subcommands(NewCommand("sub", Flags(StringFlag("flag")))),
			),
			CompareCommands,
		}, {
			"command with run function",
			&Command{Name: "sample", Flags: make(FlagsSet), Run: SampleRun, Subcommands: make(CommandsSet)},
			NewCommand("sample", RunFunction(SampleRun)),
			CompareCommands,
		}, {
			"command with runE function",
			&Command{Name: "sample", Flags: make(FlagsSet), RunE: SampleRunE, Subcommands: make(CommandsSet)},
			NewCommand("sample", RunEFunction(SampleRunE)),
			CompareCommands,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompareCommands(t, tt.expected, tt.actual)
		})
	}
}
