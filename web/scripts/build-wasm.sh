#!/usr/bin/env bash
# Builds the engine for the browser into static/: mms.wasm plus Go's
# wasm_exec.js loader, both matching the installed Go toolchain. Also
# writes the language reference the docs pages are built from.
set -euo pipefail
cd "$(dirname "$0")/../../engine"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o ../web/static/mms.wasm ./cmd/mms-wasm
root="$(go env GOROOT)"
wasm_exec="$root/lib/wasm/wasm_exec.js"
[ -f "$wasm_exec" ] || wasm_exec="$root/misc/wasm/wasm_exec.js"
# install, not cp: GOROOT is often read-only (a toolchain Go downloaded
# into the module cache, Homebrew, Nix), and cp keeps the 0444 mode, so
# the next build could not overwrite the copy. install replaces the file
# and sets a writable mode, which also mends a copy an older build left.
install -m 644 "$wasm_exec" ../web/static/wasm_exec.js
ls -l ../web/static/mms.wasm
# The language reference for the docs pages.
go run ./cmd/mms reference > ../web/src/lib/docs/reference.json
