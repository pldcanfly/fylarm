IS_NIXOS := $(shell if [ -d "/etc/NIXOS" ]; then echo true; else echo false; fi)

NIX_DEPS := libGL pkg-config xorg.libX11.dev xorg.libXcursor xorg.libXi xorg.libXinerama xorg.libXrandr xorg.libXxf86vm alsa-lib

NIX_SHELL_CMD := nix-shell -p $(NIX_DEPS) --command

build:
ifeq ($(IS_NIXOS),true)
	@echo "Running with Nix..."
	$(NIX_SHELL_CMD) "make _build"
else
	@echo "Running without Nix..."
	make _build
endif

_build:
	go build -o bin/fylarm ./cmd/fylarm/.

run:
ifeq ($(IS_NIXOS),true)
	@echo "Running with Nix..."
	$(NIX_SHELL_CMD) "make _run"
else
	@echo "Running without Nix..."
	make _run
endif

run:
	air
