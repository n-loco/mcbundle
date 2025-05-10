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
native-target-path = $(call native-package-path,$(1))$(DBGPATH)bpbuild$(call exe-suffix,$(1))

# param $(1): target list
cross-native-target-paths = $(strip $(foreach target,$(1), $(call native-target-path,$(target))))

# ========== variables =========== #

GO_SOURCES := $(call rwildcard,**/*.go)
BPBUILD_TS_SOURCES := $(call rwildcard,npm/bpbuild/source/**/*.ts)
CREATE_TS_SOURCES := $(call rwildcard,npm/create/source/**/*.ts)

GO_LINKER_FLAGS := $(if $(IS_RELEASE),-s -w,)

NATIVE_TARGET := $(call native-target-path,$(DEFAULT_PLATFORM))
CROSS_NATIVE_TARGETS := $(call cross-native-target-paths,$(ALL_PLATFORMS))
JS_TARGETS := npm/bpbuild/$(DBGPATH)dist/bpbuild.mjs npm/create/$(DBGPATH)dist/create.mjs

SOURCE_ASSETS := $(wildcard assets/*)
IMPORTED_ASSETS := $(SOURCE_ASSETS:%=internal/%.go)

# ======== helper targets ======== #

.PHONY::	build
build:	import-assets npm/pnpm-lock.yaml $(NATIVE_TARGET) $(JS_TARGETS)

.PHONY::	build-cross
build-cross:	import-assets npm/pnpm-lock.yaml $(CROSS_NATIVE_TARGETS) $(JS_TARGETS)

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
clean:	clean-imported-assets clean-native-builds clean-js-builds

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
	@$(RM) ./npm/bpbuild/debug/dist ./npm/create/debug/dist

.PHONY::	clean-node-modules
clean-node-modules:
	@$(RM) ./npm/node_modules ./npm/bpbuild/node_modules ./npm/create/node_modules

# ========= true targets ========= #

internal/assets/%.go:	assets/% $(IMPORT_ASSET_DEPS)
	@$(IMPORT_ASSET) $*
	@go fmt ./$@

# param $(1): target
define bpbuild-binary-template
npm/@bpbuild/$(1)/$(DBGPATH)bpbuild$(call exe-suffix,$(1)): export GOOS = $(call target-to-goos,$(1))
npm/@bpbuild/$(1)/$(DBGPATH)bpbuild$(call exe-suffix,$(1)): export GOARCH = $(call target-to-goarch,$(1))
npm/@bpbuild/$(1)/$(DBGPATH)bpbuild$(call exe-suffix,$(1)):	go.mod go.sum $(GO_SOURCES)
	@go build -o $$@ -ldflags="$(GO_LINKER_FLAGS)" ./main.go
endef
$(foreach platform,$(ALL_PLATFORMS),$(eval $(call bpbuild-binary-template,$(platform))))
undefine bpbuild-binary-template

npm/bpbuild/$(DBGPATH)dist/bpbuild.mjs:	npm/bpbuild/package.json npm/bpbuild/esbuild.js $(BPBUILD_TS_SOURCES)
	@cd npm && cd bpbuild && pnpm --silent build

npm/create/$(DBGPATH)dist/create.mjs:	npm/create/package.json npm/create/esbuild.js $(CREATE_TS_SOURCES)
	@cd npm && cd create && pnpm --silent build

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
