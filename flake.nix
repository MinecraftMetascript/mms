{
  description = "DevShell with Antlr4 and Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/d98abf5cf5914e5e4e9d57205e3af55ca90ffc1d";
    flake-utils.url = "github:numtide/flake-utils";

  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.antlr4
            pkgs.go
            pkgs.fish

            (pkgs.writeShellApplication{
              name = "debug-file";
              runtimeInputs = [
                pkgs.antlr4
                pkgs.go
              ];
              text = ''
                antlr4-parse grammars/MMSParser.g4 grammars/MMSLexer.g4 mmsFile -gui "$1"
              '';
                
            })

            (pkgs.writeShellApplication {
              name = "antlr-build";
              runtimeInputs = [
                pkgs.antlr4
                pkgs.go
              ];
              text = ''
                set -e
                src="./grammars"
                dst="./lang"
                ${pkgs.antlr4}/bin/antlr4 -Dlanguage=Go $src/MMSLexer.g4 -o $dst -package grammars;
                ${pkgs.antlr4}/bin/antlr4 -Dlanguage=Go $src/MMSParser.g4 -lib $dst/grammars -o $dst -package grammars;                
              '';
            })
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
                  grep Keyword_ < grammars/MMSLexer.g4 | sed -E 's/:.*/ | /' | tr -d '\n' | sed 's/ | $/\n/'
                  echo ";"
                } > grammars/MMS_Keyword_Rule.g4
              '';
            })

          ];

          shellHook = ''
            exec fish
          '';
        };

        devShells.jetbrainsPlugin = pkgs.mkShell {
          buildInputs = [ pkgs.jdk17 pkgs.gradle ];
          shellHook = ''
            echo "Use: nix develop .#jetbrainsPlugin -c gradle --no-daemon buildPlugin (run from extensions/jetbrains)"
          '';
        };

        packages.default = pkgs.buildGoModule {
          pname = "mms";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;
          # If you use vendoring, run `go mod vendor` and replace null with the hash
          subPackages = [ "." ];

          preBuild = ''
                dst=./lang
                mkdir -p $dst
                # Specifying dst/grammars is required here for some reason -- but not in the shell script?
                ${pkgs.antlr4}/bin/antlr4 -Dlanguage=Go $src/grammars/MMSLexer.g4 -o $dst/grammars -package grammars;
                ${pkgs.antlr4}/bin/antlr4 -Dlanguage=Go $src/grammars/MMSParser.g4 -lib $dst/grammars -o $dst/grammars -package grammars;                
          '';
        };

        packages.wasm = pkgs.stdenvNoCC.mkDerivation {
          pname = "mms-wasm";
          version = "0.1.0";
          src = ./.;
          buildInputs = [ pkgs.go ];
          buildPhase = ''
            export GOOS=js
            export GOARCH=wasm
            export GOCACHE="$TMPDIR/go-cache"
            go build -o main.wasm
          '';
          installPhase = ''
            mkdir -p $out
            mv main.wasm $out/
            cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" $out/
          '';
        };

        packages.jetbrainsPlugin = pkgs.stdenvNoCC.mkDerivation {
          pname = "mms-jetbrains-plugin";
          version = "0.1.0";
          src = ./extensions/jetbrains;
          nativeBuildInputs = [ pkgs.jdk17 pkgs.gradle ];
          buildPhase = ''
            set -eu
            export JAVA_HOME=${pkgs.jdk17}
            export GRADLE_USER_HOME=$TMPDIR/gradle-home
            gradle --no-daemon buildPlugin
          '';
          installPhase = ''
            set -eu
            mkdir -p $out
            cp -v build/distributions/*.zip $out/
          '';
        };
        
      }
    );
}
