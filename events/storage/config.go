package storage

var defaultConfig = Config{
	BatchSize: 100,
}

// Config allows adjusting behavior of the storage component.
type Config struct {
	// BatchSize limits the total number of events returned in a single batch.
	BatchSize uint
}

// OptionFunc can be used to adjust a config option.
type OptionFunc func(*Config)

func WithBatchSize(size uint) OptionFunc {
	return func(cfg *Config) {
		cfg.BatchSize = size
	}
}
