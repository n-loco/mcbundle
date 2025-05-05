# ========== variables =========== #

IMPORT_ASSET_DEPS := scripts/assets.py
PLATFORM_HELPER_DEPS := scripts/platform.py scripts/compat.py
UPDATE_PKGS_DEPS := scripts/update_pkgs.py scripts/compat.py

FINDEXEC := ./scripts/findexec/script.$(SHELL)
FINDEXEC := $(if $(WIN32),$(subst /,\,$(FINDEXEC)),$(FINDEXEC))

IMPORT_ASSET := $(PYTHON_RT) scripts/assets.py
UPDATE_PKGS := $(PYTHON_RT) scripts/update_pkgs.py

RM := $(PYTHON_RT) scripts/rm.py

PLATFORM_HELPER := $(PYTHON_RT) scripts/platform.py

# ========== functions =========== #

# param $(1): executable
findexec = $(if $(shell $(FINDEXEC) $(1)),,$(1))

# yes, it is just a variable
platform-wildcard = $(shell $(PLATFORM_HELPER) --platform-wildcard)

# param $(1): target
exe-suffix = $(shell $(PLATFORM_HELPER) --exe-suffix $(1))

# param $(1): target
target-to-goos = $(shell $(PLATFORM_HELPER) --node-os-to-goos $(word 1, $(subst -, ,$(1))))

# param $(1): target
target-to-goarch = $(shell $(PLATFORM_HELPER) --node-cpu-to-goarch $(word 2, $(subst -, ,$(1))))
