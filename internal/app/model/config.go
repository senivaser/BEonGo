package model

type Config struct {
	Uri      string `toml:"mongo_uri"`
	Database string `toml:"mongo_db_name"`
}

func NewConfig() *Config {
	return &Config{}
}
