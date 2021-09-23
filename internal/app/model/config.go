package model

type Config struct {
	uri      string `toml:"mongo_uri"`
	database string `toml:"mongo_db_name"`
}

func NewConfig() *Config {
	return &Config{}
}
