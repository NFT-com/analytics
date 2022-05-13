package api

var defaultConfig = Config{
	SearchLimit: 1000,
}

// Config allows customizing the server behavior.
type Config struct {
	// SearchLimit limits the number of records returned by the NFT search query.
	SearchLimit uint
}

// OptionFunc can be used to adjust a config option.
type OptionFunc func(*Config)

// WithSearchLimit sets the limit to the number of records returned by the NFT search query.
func WithSearchLimit(limit uint) OptionFunc {
	return func(cfg *Config) {
		cfg.SearchLimit = limit
	}
}
