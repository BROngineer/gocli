package cli

type FlagOption func(f Flag)

func FlagDescription(value string) FlagOption {
	return func(f Flag) {
		f.SetDescription(value)
	}
}

func Shorthand(value string) FlagOption {
	return func(f Flag) {
		f.SetShorthand(value)
	}
}

func Required() FlagOption {
	return func(f Flag) {
		f.SetRequired()
	}
}

func Shared() FlagOption {
	return func(f Flag) {
		f.SetShared()
	}
}

func Default[T allowed](value T) FlagOption {
	return func(f Flag) {
		f.setDefVal(Value[T]{defined: true, val: &value})
	}
}
