# Configurable variables:
#	PYTHON:
#		Python runtime
#	PYTHON_FLAGS:
#		Flags for python's runtime

include setup.mk

# ========== functions =========== #

# param $(1): target
native-package-path = npm/@bpbuild/$(1)/

# param $(1): target list
native-package-paths = $(strip $(foreach target,$(1), $(call native-package-path,$(target))))

# param $(1): target list
native-packagejson-paths = $(strip $(foreach target,$(1), $(call native-package-path,$(target))package.json))

# param $(1): target
target-binary-path = $(call native-package-path,$(1))bpbuild$(call exe-suffix,$(1))

# param $(1): target list
cross-binary-target-paths = $(strip $(foreach target,$(1), $(call target-binary-path,$(target))))

# ========== variables =========== #

GO_SOURCES := $(call rwildcard,**/*.go)

TARGET_BINARY := $(call target-binary-path,$(DEFAULT_PLATFORM))
CROSS_BINARY_TARGETS := $(call cross-binary-target-paths,$(ALL_PLATFORMS))

SOURCE_ASSETS := $(wildcard assets/*)
IMPORTED_ASSETS := $(SOURCE_ASSETS:%=internal/%.go)

# ======== helper targets ======== #

.PHONY::	build
build:	import-assets npm/pnpm-lock.yaml $(TARGET_BINARY)
	@cd npm && cd bpbuild && pnpm --silent build
	@cd npm && cd create && pnpm --silent build

.PHONY::	build-cross
build-cross:	import-assets npm/pnpm-lock.yaml $(CROSS_BINARY_TARGETS)
	@cd npm && cd bpbuild && pnpm --silent build
	@cd npm && cd create && pnpm --silent build

.PHONY::	import-assets
import-assets:	$(IMPORTED_ASSETS)

.PHONY::	fmt
fmt:
	@go fmt ./...

.PHONY::	setup
setup:	import-assets npm/pnpm-lock.yaml
	@go get golang.org/x/sys@v0.31.0 github.com/evanw/esbuild@v0.25.2
	@cd npm && pnpm --silent install

.PHONY::	clean
clean:	clean-imported-assets clean-native-builds clean-js-builds clean-node-modules

.PHONY::	clean-imported-assets
clean-imported-assets:
	@$(RM) ./internal/assets

.PHONY::	clean-native-builds
clean-native-builds:
	@$(RM) --unused-builds
	@$(RM) $(call native-package-paths,$(ALL_PLATFORMS))

.PHONY::	clean-js-builds
clean-js-builds:
	@$(RM) ./npm/bpbuild/dist ./npm/create/dist

.PHONY::	clean-node-modules
clean-node-modules:
	@$(RM) ./npm/node_modules ./npm/bpbuild/node_modules ./npm/create/node_modules

# ========= true targets ========= #

internal/assets/%.go:	assets/% $(IMPORT_ASSET_DEPS)
	@$(IMPORT_ASSET) $*
	@go fmt ./$@

# param $(1): target
define bpbuild-binary-template
npm/@bpbuild/$(1)/bpbuild$(call exe-suffix,$(1)):
  export GOOS = $(call target-to-goos,$(1))
  export GOARCH = $(call target-to-goarch,$(1))
npm/@bpbuild/$(1)/bpbuild$(call exe-suffix,$(1)):	go.mod go.sum $(GO_SOURCES)
	@go build -o $$@ ./main.go
endef
$(foreach platform,$(ALL_PLATFORMS),$(eval $(call bpbuild-binary-template,$(platform))))
undefine bpbuild-binary-template

# native package.json
npm/@bpbuild/%/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --target $*

npm/bpbuild/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --main-package

npm/create/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --create-package

npm/pnpm-lock.yaml::
  NATIVE_PACKAGEJSON_PATHS := $(call native-packagejson-paths,$(ALL_PLATFORMS))
npm/pnpm-lock.yaml::	$(UPDATE_PKGS_DEPS) npm/pnpm-workspace.yaml npm/package.json
npm/pnpm-lock.yaml::	$(NATIVE_PACKAGEJSON_PATHS) npm/bpbuild/package.json npm/create/package.json
	@$(RM) --unused-builds
	@cd npm && pnpm --silent install --lockfile-only
