# [WIP!!!] Neogoon

Gooning encouragement utility.

# Differences from Edgeware

- It functions in any capacity
- Configuration is mostly through .toml, see `example/neogoon.toml` for a well-documented example.
- The package (set) format is different because Edgeware's isn't very good nor well documented.
- Cross-platform (Linux + macOS support)
- Portable compiled binary = much simpler to set up (and it's faster I guess)
- Walltaker integration
- New exclusive annoyances!

# Build Instructions

Just `go build`; on Windows you need GCC -- try https://jmeubank.github.io/tdm-gcc/, and build
with `build_windows.bat` or it'll spawn a command prompt along with the UI.