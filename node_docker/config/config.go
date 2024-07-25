package config

type Config struct {
}

func LoadConfig() (*Config, error) {
	// Create a new Config instance
	config := &Config{}

	return config, nil
}
