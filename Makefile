prepare: 
	rm -rf ./build/*

build: prepare
	go build -o build/fylarm main.go

run: 
	air
	# nix-shell -p pkg-config alsa-lib --command air
