package gocli

type Config interface {
	LoadFromFiles(paths []string) error
	LoadFromEnv(prefix string) error
	LoadFromFlags(flags FlagSet) error
}

func TypedConfig[T any](cfg Config) *T {
	return any(cfg).(*T)
}
