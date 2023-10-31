package gocli

type FlagOption func(f flagDataSetter)

func FlagDescription(value string) FlagOption {
	return func(f flagDataSetter) {
		f.setDescription(value)
	}
}

func Shorthand(value string) FlagOption {
	return func(f flagDataSetter) {
		f.setShorthand(value)
	}
}

func Required() FlagOption {
	return func(f flagDataSetter) {
		f.setRequired()
	}
}

func Shared() FlagOption {
	return func(f flagDataSetter) {
		f.setShared()
	}
}

func Default[T allowed](value T) FlagOption {
	return func(f flagDataSetter) {
		f.setDefVal(Value[T]{defined: true, val: &value})
	}
}
