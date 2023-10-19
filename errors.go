package gocli

import (
	"errors"
	"fmt"
)

const (
	IntParseErrorMessage           = "failed to parse input as int"
	BoolParseErrorMessage          = "failed to parse input as bool"
	DurationParseErrorMessage      = "failed to parse input as time.Duration"
	MissedDefaultValueErrorMessage = "default value is mandatory for optional flag %s"
	MissedRequiredFlagErrorMessage = "flag %s is required"
	MissedFlagValueErrorMessage    = "value required for flag %s"
	UndefinedFlagErrorMessage      = "flag -%s is not defined for command %s"
	UndefinedCommandErrorMessage   = "undefined command %s"
	FlagNotFoundErrorMessage       = "failed to get flag"
	FlagTypeMismatchErrorMessage   = "failed to get flag value"
	CommandRunFailed               = "failed to execute command"
)

var (
	flagNotFoundError     = errors.New("flag with provided name does not exist in current FlagSet")
	flagTypeMismatchError = errors.New("stored value type differs to requested type")
	commandExecuteError   = errors.New("no function to run defined")
)

type FlagError struct {
	message string
	inner   error
}

func (e FlagError) Error() string {
	if e.inner == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %s", e.message, e.inner.Error())
}

func (e FlagError) Is(target error) bool {
	var err FlagError
	if ok := errors.As(target, &err); ok {
		return e.message == target.(FlagError).message
	}
	return false
}

func (e FlagError) Unwrap() error {
	return e.inner
}

func (e FlagError) Wrap(err error) error {
	e.inner = err
	return e
}

func NewFlagError(msg string) FlagError {
	return FlagError{message: msg}
}

func ParseIntError() FlagError {
	return NewFlagError(IntParseErrorMessage)
}

func ParseBoolError() FlagError {
	return NewFlagError(BoolParseErrorMessage)
}

func ParseDurationError() FlagError {
	return NewFlagError(DurationParseErrorMessage)
}

type FlagSetError struct {
	message string
	inner   error
}

func (e FlagSetError) Error() string {
	if e.inner == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %s", e.message, e.inner.Error())
}

func (e FlagSetError) Unwrap() error {
	return e.inner
}

func (e FlagSetError) Wrap(err error) error {
	e.inner = err
	return e
}

func (e FlagSetError) Is(target error) bool {
	var err FlagSetError
	if ok := errors.As(target, &err); ok {
		return e.message == target.(FlagSetError).message
	}
	return false
}

func NewFlagSetError(msg string) FlagSetError {
	return FlagSetError{message: msg}
}

func FlagNotFoundError() error {
	return NewFlagSetError(FlagNotFoundErrorMessage).Wrap(flagNotFoundError)
}

func FlagTypeMismatchError() error {
	return NewFlagSetError(FlagTypeMismatchErrorMessage).Wrap(flagTypeMismatchError)
}

type CommandError struct {
	message string
	inner   error
}

func (e CommandError) Error() string {
	if e.inner == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %s", e.message, e.inner.Error())
}

func (e CommandError) Unwrap() error {
	return e.inner
}

func (e CommandError) Wrap(err error) error {
	e.inner = err
	return e
}

func (e CommandError) Is(target error) bool {
	var err CommandError
	if ok := errors.As(target, &err); ok {
		return e.message == target.(CommandError).message
	}
	return false
}

func NewCommandError(msg string) CommandError {
	return CommandError{message: msg}
}

func CommandExecuteError() error {
	return NewCommandError(CommandRunFailed).Wrap(commandExecuteError)
}
