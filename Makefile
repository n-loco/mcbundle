# ============ config ============ #

include setup.mk

# ========== functions =========== #

# param $(1): generic string
# returns .exe if $(1) has win32
# otherwise returns empty
exename = $(if $(findstring win32,$(1)),.exe)

# ========== variables =========== #

# matches the pattern `${process.platform}-${process.arch}`
HOST_PLATFORM := $(shell node -e "process.stdout.write(process.platform+'-'+process.arch)")

# the items in the list MUST match the pattern `${process.platform}-${process.arch}`
SUPPORTED_PLATFORMS := $(sort $(file < platforms.txt))
export SUPPORTED_PLATFORMS

# matches the pattern `${process.platform}-${process.arch}`
# it is empty if the host platform is not supported as a target
DEFAULT_PLATFORM := $(if $(filter $(HOST_PLATFORM),$(SUPPORTED_PLATFORMS)),$(HOST_PLATFORM))

GO_LINKER_FLAGS := $(if $(IS_RELEASE),-s -w)

MCBUNDLE_GO_SOURCES := main.go $(call glob,internal/**/*.go)
MCBUNDLE_TS_SOURCES := $(call glob,npm/mcbundle/source/**/*.ts)
MCBUNDLE_ASSETS := $(wildcard assets/*)
CREATE_TS_SOURCES := $(call glob,npm/create/source/**/*.ts)

JS_LIB_TOOLS_JS := $(call glob,npm/tools/lib/**/*.js)
JS_LIB_TOOLS_DTS := $(JS_LIB_TOOLS_JS:%.js=%.d.ts)
JS_LIB_TOOLS_FILES := $(JS_LIB_TOOLS_JS) $(JS_LIB_TOOLS_DTS)

DIST_DIR := $(if $(IS_DEBUG),debug,dist)

NATIVE_PACKAGES := $(SUPPORTED_PLATFORMS:%=npm/@mcbundle/%)
NATIVE_PACKAGE_JSONS := $(NATIVE_PACKAGES:%=%/package.json)

GHOST_NATIVE_PACKAGES := $(filter-out $(NATIVE_PACKAGES),$(wildcard npm/@mcbundle/*))

DEFAULT_NATIVE_TARGET := npm/@mcbundle/$(DEFAULT_PLATFORM)/$(DIST_DIR)/mcbundle$(call exename,$(DEFAULT_PLATFORM))
DEFAULT_NATIVE_TARGET := $(if $(DEFAULT_PLATFORM),$(DEFAULT_NATIVE_TARGET))

CROSS_NATIVE_TARGETS := $(foreach pkg,$(NATIVE_PACKAGES),$(pkg)/$(DIST_DIR)/mcbundle$(call exename,$(pkg)))

JS_TARGETS := npm/mcbundle/$(DIST_DIR)/mcbundle.mjs npm/create/$(DIST_DIR)/create.mjs

# ======== helper targets ======== #

.PHONY::	build
build:	npm/pnpm-lock.yaml $(DEFAULT_NATIVE_TARGET) $(JS_TARGETS)

.PHONY::	build-cross
build-cross:	npm/pnpm-lock.yaml $(CROSS_NATIVE_TARGETS) $(JS_TARGETS)

.PHONY::	fmt
fmt:
	@go fmt ./...

.PHONY::	setup
setup:	$(NATIVE_PACKAGE_JSONS)
	@go get golang.org/x/sys@v0.31.0 github.com/evanw/esbuild@v0.25.2
	@cd npm && pnpm --silent install && pnpm --silent tool-types

.PHONY::	clean
clean:	clean-native-builds clean-js-builds

.PHONY::	clean-native-builds
clean-native-builds:
	@$(RM) $(NATIVE_PACKAGES) $(GHOST_NATIVE_PACKAGES)

.PHONY::	clean-js-builds
clean-js-builds:
	@$(RM) ./npm/mcbundle/dist ./npm/create/dist
	@$(RM) ./npm/mcbundle/debug ./npm/create/debug

.PHONY::	clean-node-modules
clean-node-modules:
	@$(RM) ./npm/node_modules ./npm/mcbundle/node_modules ./npm/create/node_modules

.PHONY:: clean-ghost-native-builds
clean-ghost-native-builds:
	@$(RM) $(GHOST_NATIVE_PACKAGES)

# ========= true targets ========= #

# param $(1): target
define mcbundle-binary-template
npm/@mcbundle/$(1)/$(DIST_DIR)/mcbundle$(call exename,$(1)): GO_PAIR = $(call node2go-pair,$(1))
npm/@mcbundle/$(1)/$(DIST_DIR)/mcbundle$(call exename,$(1)): export GOOS = $$(word 1,$$(GO_PAIR))
npm/@mcbundle/$(1)/$(DIST_DIR)/mcbundle$(call exename,$(1)): export GOARCH = $$(word 2,$$(GO_PAIR))
npm/@mcbundle/$(1)/$(DIST_DIR)/mcbundle$(call exename,$(1)):	go.mod go.sum $$(MCBUNDLE_GO_SOURCES) $$(MCBUNDLE_ASSETS)
	@go build -o $$@ -ldflags="$$(GO_LINKER_FLAGS)" ./main.go
endef
$(foreach platform,$(SUPPORTED_PLATFORMS),$(eval $(call mcbundle-binary-template,$(platform))))
undefine mcbundle-binary-template

npm/mcbundle/$(DIST_DIR)/mcbundle.mjs:	$(JS_LIB_TOOLS_FILES) npm/mcbundle/esbuild.js $(MCBUNDLE_TS_SOURCES)
	@cd npm && cd mcbundle && pnpm --silent build

npm/create/$(DIST_DIR)/create.mjs:	$(JS_LIB_TOOLS_FILES) npm/create/esbuild.js $(CREATE_TS_SOURCES)
	@cd npm && cd create && pnpm --silent build

$(JS_LIB_TOOLS_DTS) &:	$(JS_LIB_TOOLS_JS)
	@cd npm && pnpm --silent tool-types

$(NATIVE_PACKAGE_JSONS) &:	$(NPKGS_DEPS) platforms.txt
	@$(NPKGS) npm/@mcbundle $(SUPPORTED_PLATFORMS)

npm/mcbundle/package.json:	npm/mcbundle/update_native.js $(NPKGS_DEPS) platforms.txt
	@cd npm && cd mcbundle && pnpm --silent sync-package

npm/create/package.json:	assets/program_version.txt
	@cd npm && cd create && pnpm --silent sync-package

npm/pnpm-lock.yaml:	platforms.txt npm/pnpm-workspace.yaml npm/package.json $(JS_LIB_TOOLS_DTS)
npm/pnpm-lock.yaml:	$(NATIVE_PACKAGE_JSONS) npm/mcbundle/package.json npm/create/package.json
	@$(RM) $(GHOST_NATIVE_PACKAGES)
	@cd npm && pnpm --silent install && pnpm --silent touch
