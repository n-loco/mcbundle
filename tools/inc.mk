# =========== commands =========== #

RM := go run ./tools/rm

NPKGS_DEPS := $(wildcard tools/npkgs/*) assets/program_version.txt assets/assets.go
NPKGS := go run ./tools/npkgs

# === function implementations === #

__GLOB := go run ./tools/glob/cmd

__NODE2GO := go run ./tools/node2go/cmd

# ========== functions =========== #

# param $(1): file pattern (or a list of patterns)
# returns a list matching files (or empty)
glob = $(shell $(__GLOB) $(1))

# param $(1): node target (MUST match the pattern `${process.platform}-${process.arch}`)
# returns equivalents $GOOS and $GOARCH respectively
node2go-pair = $(shell $(__NODE2GO) $(1))
