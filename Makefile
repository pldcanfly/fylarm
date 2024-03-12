build: 
	go build -o bin/fylarm ./cmd/fylarm/. 

run: 
	air
	# nix-shell -p pkg-config alsa-lib --command air
