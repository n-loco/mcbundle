# ============ config ============ #

ifeq ($(OS),Windows_NT)
  NULL := NUL
  SHELL := cmd
  WIN32 := 1
else
  NULL := /dev/null
  SHELL := sh
endif

# param $(1): executable name/path
# returns 1 on success or empty on failure
exec-exists = $(strip $(if $(WIN32),\
  $(shell WHERE $(1) > NUL 2> NUL && ECHO 1 & EXIT /B 0),\
  $(shell command -v $(1) >/dev/null 2>&1 && echo 1; exit 0)\
))

-include config.mk

BUILD_MODE ?= debug

ifeq ($(BUILD_MODE),release)
else ifeq ($(BUILD_MODE),debug)
else
$(error Unknown build mode: $(BUILD_MODE))
endif

export BUILD_MODE

IS_DEBUG := $(filter debug,$(BUILD_MODE))
IS_RELEASE := $(filter release,$(BUILD_MODE))

include tools/inc.mk

# ==== external dependencies ===== #

# param $(1): executable name/path
# returns $(1) if it does not exist
return-if-missing = $(if $(call exec-exists,$(1)),,$(1))

NOT_FOUND_PROGRAMS :=

NOT_FOUND_PROGRAMS += $(call return-if-missing,go)
NOT_FOUND_PROGRAMS += $(call return-if-missing,node)
NOT_FOUND_PROGRAMS += $(call return-if-missing,pnpm)

ifneq (,$(NOT_FOUND_PROGRAMS))
  $(error Programs not found: $(NOT_FOUND_PROGRAMS))
endif

undefine NOT_FOUND_PROGRAMS
undefine return-if-missing
