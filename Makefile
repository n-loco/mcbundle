# Env variables:
#	PYTHON_RT:
#		Python runtime
#	PYTHON_RT_FLAGS:
#		Flags for python's runtime

ifeq ($(PYTHON_RT),)
PYTHON_RT = python3
endif

fmt:
	@go fmt ./...

gen-assets:
	@$(PYTHON_RT) $(PYTHON_RT_FLAGS) python/assets.py
	@go fmt ./internal/assets/...

setup:	gen-assets
	@go get golang.org/x/sys@v0.31.0
	@go get github.com/evanw/esbuild@v0.25.2
