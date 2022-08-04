package config

import (
	"github.com/BurntSushi/toml"
)

// Creating mappings between "idiomatic" TOML key naming and
// idiomatic Go struct names is a lot of effort, but improves
// developer and user experience, so I'm bothering.

type Config struct {
	Infection   Infection   `toml:"infection"`
	Hiberate    Hibernate   `toml:"hibernate"`
	Annoyances  Annoyances  `toml:"annoyances"`
	DriveFiller DriveFiller `toml:"drive-filler"`
	Downloader  Downloader  `toml:"downloader"`
}

type Infection struct {
	StartOnBoot bool `toml:"start-on-boot"`
	HideSystray bool `toml:"hide-systray"`
}

type Hibernate struct {
	Enabled      bool `toml:"enabled"`
	MinimumWait  int  `toml:"minimum-wait"`
	MaximumWait  int  `toml:"maximum-wait"`
	ActivityTime int  `toml:"activity-time"`
}

type Downloader struct {
	Enabled      bool     `toml:"enabled"`
	Booru        string   `toml:"booru"`
	Tags         []string `toml:"tags"`
	MinimumScore int      `toml:"minimum-score"`
}

func Load(c *Config, filename string) error {
	_, err := toml.DecodeFile(filename, &c)
	if err != nil {
		return err
	}

	return err
}
