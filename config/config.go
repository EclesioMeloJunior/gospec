package config

import "flag"

// Config gospec definition
type Config struct {
	ApispecFilesFlag string
}

func loadFlags(config *Config) {
	flag.StringVar(&config.ApispecFilesFlag, "test-files", "", "Specify the relative path to test files")

	flag.Parse()
}

// Load will get the app entrypoint
func Load() *Config {
	config := &Config{}

	loadFlags(config)

	return config
}
