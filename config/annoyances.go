package config

type Annoyances struct {
	Enabled        bool `toml:"enabled"`
	SecondsPerTick int  `toml:"seconds-per-tick"`

	// Annoyance configuration
	Popups             Popups             `toml:"popups"`
	Prompts            Prompts            `toml:"prompts"`
	Audio              Audio              `toml:"audio"`
	AutoType           AutoType           `toml:"auto-type"`
	OverwriteClipboard OverwriteClipboard `toml:"overwrite-clipboard"`
}

type Popups struct {
	Chance             int      `toml:"chance"`
	Media              []string `toml:"media"`
	Opacity            int      `toml:"opacity"`
	DenialChance       int      `toml:"denial-chance"`
	Timeout            int      `toml:"timeout"`
	AllowManualClosing bool     `toml:"allow-manual-closing"`
	Volume             int      `toml:"volume"`
	Mitosis            Mitosis  `toml:"mitosis"`
}

type Mitosis struct {
	Strength           int  `toml:"strength"`
	TriggeredByTimeout bool `toml:"triggered-by-timeout"`
}

type Prompts struct {
	Chance      int `toml:"chance"`
	MaxMistakes int `toml:"max-mistakes"`
}

type Audio struct {
	Chance      int `toml:"chance"`
	MaxPlaytime int `toml:"max-playtime"`
	Volume      int `toml:"volume"`
}

type AutoType struct {
	Chance     int  `toml:"chance"`
	PressEnter bool `toml:"press-enter"`
}

type OverwriteClipboard struct {
	Chance int `toml:"chance"`
}
