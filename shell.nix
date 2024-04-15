{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  packages = with pkgs; [ 
	libGL
	pkg-config
	xorg.libX11.dev
	xorg.libXcursor
	xorg.libXi
	xorg.libXinerama
	xorg.libXrandr
	xorg.libXxf86vm
	alsa-lib
  ];
}


