package config

import (
	"path/filepath"

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
	Wallpaper   Wallpaper   `toml:"wallpaper"`
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

type DriveFiller struct {
	Enabled    bool       `toml:"enabled"`
	Rate       int        `toml:"rate"`
	Root       string     `toml:"root"`
	Filenames  []string   `toml:"filenames"`
	Downloader Downloader `toml:"downloader"`
}

type Downloader struct {
	Booru        string   `toml:"booru"`
	Tags         []string `toml:"tags"`
	MinimumScore int      `toml:"minimum-score"`
}

type Wallpaper struct {
	Enabled   bool   `toml:"enabled"`
	Rate      int    `toml:"rate"`
	Mode      string `toml:"mode"`
	PreferFit bool   `toml:"prefer-fitting-ratio"`
}

func Load(c *Config, filename string) error {
	_, err := toml.DecodeFile(filename, &c)
	if err != nil {
		return err
	}

	// Normalise paths
	c.DriveFiller.Root = filepath.FromSlash(c.DriveFiller.Root)

	return err
}
