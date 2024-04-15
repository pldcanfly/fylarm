HAS_NIX_SHELL := $(shell command -v nix-shell 2> /dev/null)

ifeq ($(strip $(HAS_NIX_SHELL)),)
    $(error "nix-shell is not installed. Please install Nix or ensure it's in your PATH.")
endif

build:
	nix-shell --command "go build -o bin/fylarm ./cmd/fylarm/."

run:
	nix-shell --command air
