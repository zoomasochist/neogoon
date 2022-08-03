# [WIP] Neogoon

Gooning encouragement utility.

# Differences from Edgeware (apart from the featureset)

- Configuration is mostly through .toml, see `example/neogoon.toml` for a well-documented example.
- The package (set) format is different because Edgeware's isn't very good. Mostly for technical reasons, but it's also not well documented (or designed). See `set/set.go` for more, if you're interested.

# Features

- Implements all of Edgeware's features
- Cross-platform
- Portable compiled binary = much simpler to set up (and it's faster I guess)
- Walltaker integration
- New exclusive annoyances!

# Build Instructions

Just `go build`; on Windows you need GCC -- try https://jmeubank.github.io/tdm-gcc/.