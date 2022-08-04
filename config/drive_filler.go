package config

type DriveFiller struct {
	Enabled   bool     `toml:"enabled"`
	Rate      int      `toml:"rate"`
	Root      string   `toml:"root"`
	Filenames []string `toml:"filenames"`
}
