# shell.nix
let
  pkgs = import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/cf8cc1201be8bc71b7cbbbdaf349b22f4f99c7ae.tar.gz") {};
in pkgs.mkShell {
    nativeBuildInputs = with pkgs; [
        go
        nodejs
    ];

    shellHook=''
        echo "Welcome to developer environment!"
        alias generate="./scripts/generate.sh"
        alias run="./scripts/run.sh"
        alias nix-web="./scripts/run_web.sh"
        alias nix-tunnel="./scripts/tunnel.sh"
    '';
}
