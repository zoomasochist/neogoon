package config

// Creating mappings between "idiomatic" TOML key naming and
// idiomatic Go struct names is a lot of effort, but improves
// developer and user experience, so I'm bothering.

type Config struct {
	Infection   Infection   `toml:"infection"`
	Hiberate    Hibernate   `toml:"hibernate"`
	Annoyances  Annoyances  `toml:"annoyances"`
	DriveFiller DriveFiller `toml:"drive-filler"`
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
