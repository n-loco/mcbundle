# Configurable variables:
#	PYTHON:
#		Python runtime
#	PYTHON_FLAGS:
#		Flags for python's runtime

include setup.mk

# ========== variables =========== #

GO_SOURCES := $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./... 2> $(NULL))
SOURCE_ASSETS := $(wildcard assets/*)

DEFAULT_EXECUTABLE_TARGET := npm/@bpbuild/$(DEFAULT_PLATFORM)/

ALL_EXECUTABLE_TARGETS := $(foreach p,$(ALL_PLATFORMS),npm/@bpbuild/$(p)/)
ALL_EXECUTABLE_TARGETS_PACKAGES := $(foreach p,$(ALL_EXECUTABLE_TARGETS),$(p)package.json)

TARGET_ASSETS := $(foreach a,$(SOURCE_ASSETS),internal/$(a).go)

# ======== helper targets ======== #

.PHONY::	build-default
build-default:	gen-assets $(DEFAULT_EXECUTABLE_TARGET) update-packages
	@cd npm && cd bpbuild && pnpm build > $(NULL)
	@cd npm && cd create && pnpm build > $(NULL)

.PHONY::	build-all
build-all:	gen-assets $(ALL_EXECUTABLE_TARGETS) update-packages
	@cd npm && cd bpbuild && pnpm build > $(NULL)
	@cd npm && cd create && pnpm build > $(NULL)

.PHONY::	update-packages
update-packages:	clean-ghost-builds $(ALL_EXECUTABLE_TARGETS_PACKAGES) npm/bpbuild/package.json npm/create/package.json
	@cd npm && pnpm install > $(NULL)

.PHONY::	gen-assets
gen-assets:	$(TARGET_ASSETS)

.PHONY::	fmt
fmt:
	@go fmt ./...

.PHONY::	setup
setup:	gen-assets
	@go get	\
	golang.org/x/sys@v0.31.0	\
	github.com/evanw/esbuild@v0.25.2

.PHONY::	clean
clean:	clean-assets clean-builds clean-builds-js clean-node-modules clean-ghost-builds

.PHONY::	clean-assets
clean-assets:
	@$(RM) ./internal/assets

.PHONY::	clean-builds
clean-builds:
	@$(RM) $(ALL_EXECUTABLE_TARGETS)

.PHONY::	clean-build-js
clean-builds-js:
	@$(RM) ./npm/bpbuild/dist ./npm/create/dist

.PHONY::	clean-node-modules
clean-node-modules:
	@$(RM) ./npm/node_modules ./npm/bpbuild/node_modules ./npm/create/node_modules

.PHONY::	clean-ghost-builds
clean-ghost-builds:
	@$(RM) --unused-builds

# ========= true targets ========= #

internal/assets/%.go:	assets/% $(GEN_ASSET_DEPS)
	@$(GEN_ASSET) $*
	@go fmt ./$@

# % here must follow the format "{os}-{arch}"
# {os} and {arch} corresponds to `process.platform` and `process.arch` in Node.js respectively
# examples: win32-x64, linux-ia32, darwin-arm64, etc
npm/@bpbuild/%/:	$(GO_SOURCES) $(PLATFORM_HELPER_DEPS)
	@GOOS=$(call target-to-goos,$*)	\
	GOARCH=$(call target-to-goarch,$*)	\
	go build -o $@ ./...

# the same rules as above apply here
npm/@bpbuild/%/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --target $*

npm/bpbuild/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --main-package

npm/create/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --create-package
