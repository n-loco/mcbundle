# Env variables:
#	PYTHON:
#		Python runtime
#	PYTHON_FLAGS:
#		Flags for python's runtime

ifeq ($(shell go env GOHOSTOS),windows)
NULL = NUL
SHELL := cmd
else
NULL = /dev/null
SHELL := sh
endif

GO_SOURCES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./... 2> $(NULL))
SOURCE_ASSETS = $(wildcard assets/*)

PYTHON ?= python3

GEN_ASSET_DEPS = python/assets.py
PLATFORM_HELPER_DEPS = python/platform.py python/compat.py
UPDATE_PKGS_DEPS = python/update_pkgs.py python/compat.py

GEN_ASSET = $(PYTHON) $(PYTHON_FLAGS) python/assets.py
PLATFORM_HELPER = $(PYTHON) $(PYTHON_FLAGS) python/platform.py
UPDATE_PKGS = $(PYTHON) $(PYTHON_FLAGS) python/update_pkgs.py
RM = $(PYTHON) $(PYTHON_FLAGS) python/rm.py

PLATFORMS = $(shell $(PLATFORM_HELPER) --platform-wildcard)

TARGET_PLATFORMS = $(foreach p,$(PLATFORMS),npm/@bpbuild/$(p)/)
TARGET_PLATFORMS_PACKAGES = $(foreach p,$(TARGET_PLATFORMS),$(p)package.json)
TARGET_ASSETS = $(foreach a,$(SOURCE_ASSETS),internal/$(a).go)

CLEAN_BUILDS = $(foreach b,$(PLATFORMS),clean-build-$(b))

build:	update-packages gen-assets $(TARGET_PLATFORMS)
	@cd npm && cd bpbuild && pnpm build > $(NULL)
	@cd npm && cd create && pnpm build > $(NULL)

update-packages:	clean-unused-builds $(TARGET_PLATFORMS_PACKAGES) npm/bpbuild/package.json npm/create/package.json
	@cd npm && pnpm install > $(NULL)

gen-assets:	$(TARGET_ASSETS)

fmt:
	@go fmt ./...

setup:	gen-assets
	@go get	\
	golang.org/x/sys@v0.31.0	\
	github.com/evanw/esbuild@v0.25.2

clean:	clean-assets clean-builds clean-build-js clean-unused-builds

clean-assets:
	@$(RM) ./internal/assets

clean-builds:	$(CLEAN_BUILDS)

clean-build-js:
	@$(RM) ./npm/bpbuild/dist
	@$(RM) ./npm/create/dist

clean-unused-builds:
	@$(RM) --unused-builds

# ========================== #

internal/assets/%.go:	assets/% $(GEN_ASSET_DEPS)
	@$(GEN_ASSET) $*
	@go fmt ./$@

npm/@bpbuild/%/:	$(GO_SOURCES) $(PLATFORM_HELPER_DEPS)
	@GOOS=$(shell $(PLATFORM_HELPER) --node-os-to-goos $(word 1, $(subst -, ,$*)))	\
	GOARCH=$(shell $(PLATFORM_HELPER) --node-cpu-to-goarch $(word 2, $(subst -, ,$*)))	\
	go build -o $@ ./...

npm/@bpbuild/%/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --executable $(word 1, $(subst -, ,$*)) $(word 2, $(subst -, ,$*))

npm/bpbuild/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --main-package

npm/create/package.json:	assets/program_version.txt $(UPDATE_PKGS_DEPS)
	@$(UPDATE_PKGS) --create-package

clean-build-%:
	@$(RM) ./npm/@bpbuild/$*