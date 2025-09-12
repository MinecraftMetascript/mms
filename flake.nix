{
  description = "DevShell with Antlr4 and Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/648f70160c03151bc2121d179291337ad6bc564b";
    flake-utils.url = "github:numtide/flake-utils";

  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        antlrBuild = pkgs.writeShellApplication {
          name = "antlr-build";
          runtimeInputs = [
            pkgs.antlr4
            pkgs.go
          ];
          text = ''
            set -e
            src="./grammar"
            dst="./lang"
            ${pkgs.antlr4}/bin/antlr4 -Dlanguage=Go $src/MinecraftMetascript.g4 -o $dst -package grammar;
          '';
        };
      in
      {
        devShells.default = pkgs.mkShell {
          env = {
            GOROOT = "${pkgs.go}/share/go";
          };

          buildInputs = [
            pkgs.antlr4
            pkgs.go
            pkgs.fish
            antlrBuild
            pkgs.cobra-cli
            (pkgs.writeShellApplication {
              name = "antlr-build-keyword-rule";
              runtimeInputs = [
                pkgs.antlr4
                pkgs.go
              ];
              text = ''
                set -e
                {
                  echo "parser grammar MMS_Keyword_Rule;"
                  echo "options { tokenVocab = MMSLexer; }"
                  echo "keyword: "
                  grep Keyword_ < grammars/Main_Lexer.g4 | sed -E 's/:.*/ | /' | tr -d '\n' | sed 's/ | $/\n/'
                  echo ";"
                } > grammars/MMS_Keyword_Rule.g4
              '';
            })

          ];

          shellHook = ''
            exec fish
          '';
        };

        packages.default = pkgs.buildGoModule {
          pname = "mms";
          version = "0.3.3";
          src = ./.;
          vendorHash = null;
          # If you use vendoring, run `go mod vendor` and replace null with the hash
          subPackages = [ "." ];

          preBuild = ''
            ${antlrBuild}/bin/antlr-build
          '';
        };

        packages.wasm = pkgs.stdenvNoCC.mkDerivation {
          pname = "mms-wasm";
          version = "0.3.3";
          src = ./.;
          buildInputs = [ pkgs.go ];
          buildPhase = ''
            export GOOS=js
            export GOARCH=wasm
            export HOME="$TMPDIR"
            export GOPATH="$TMPDIR/go"
            export GOMODCACHE="$GOPATH/pkg/mod"
            export GOCACHE="$TMPDIR/go-cache"
            export GOPROXY=off

            go build -o main.wasm -mod=vendor
          '';
          installPhase = ''
            mkdir -p $out/js/dist
            mv main.wasm $out/js/dist
            cp $src/wasm/* $out/js -r
            cp $src/wasm/.* $out/js -r

            cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" $out/
          '';
        };
      }
    );
}
