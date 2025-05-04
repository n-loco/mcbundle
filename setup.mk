# ============ config ============ #

ifeq ($(OS),Windows_NT)
  NULL := NUL
  SHELL := cmd
  WIN32 := 1
else
  NULL := /dev/null
  SHELL := sh
endif

-include config.mk

PYTHON ?= python3
PYTHON_RT := $(PYTHON) $(PYTHON_FLAGS)

include scripts/inc.mk

# ==== external dependencies ===== #

NOT_FOUND_PROGRAMS :=

NOT_FOUND_PROGRAMS += $(call findexec,go)
NOT_FOUND_PROGRAMS += $(call findexec,$(PYTHON))
NOT_FOUND_PROGRAMS += $(call findexec,node)
NOT_FOUND_PROGRAMS += $(call findexec,pnpm)

ifneq (,$(NOT_FOUND_PROGRAMS))
  $(error Programs not found: $(NOT_FOUND_PROGRAMS))
endif

undefine NOT_FOUND_PROGRAMS

# ========== variables =========== #

DEFAULT_PLATFORM := $(shell node -e "process.stdout.write(process.platform + '-' + process.arch)")
ALL_PLATFORMS := $(platform-wildcard)
