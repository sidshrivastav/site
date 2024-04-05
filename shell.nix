with import <nixpkgs> {};

stdenv.mkDerivation {
    name = "site";
    buildInputs = [
        go
        google-cloud-sdk
    ];
}
