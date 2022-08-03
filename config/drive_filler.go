package config

type DriveFiller struct {
	Enabled    bool       `toml:"enabled"`
	Rate       int        `toml:"rate"`
	Root       string     `toml:"root"`
	Downloader Downloader `toml:"downloader"`
}

type Downloader struct {
	Enabled      bool     `toml:"enabled"`
	Booru        string   `toml:"booru"`
	Tags         []string `toml:"tags"`
	MinimumScore int      `toml:"minimum-score"`
}
