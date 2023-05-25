{ lib, buildPackages ? import <nixpkgs> { } }:

let
  version = "1.0.0";
in buildPackages.stdenv.mkDerivation {
  name = "yock-${version}";

  src = ./.;

  meta = with lib; {
    description = "Yock is a solution of cross platform to compose distributed build stream.";
    homepage = "https://github.com/Ansurfen/yock";
    license = licenses.mit;
    platforms = [ "x86_64-linux" ];
    maintainers = [ maintainers.ansurfen ];
  };
}
