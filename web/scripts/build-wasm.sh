#!/usr/bin/env bash
# Builds the engine for the browser into static/: mms.wasm plus Go's
# wasm_exec.js loader, both matching the installed Go toolchain.
set -euo pipefail
cd "$(dirname "$0")/../../engine"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o ../web/static/mms.wasm ./cmd/mms-wasm
root="$(go env GOROOT)"
if [ -f "$root/lib/wasm/wasm_exec.js" ]; then
	cp "$root/lib/wasm/wasm_exec.js" ../web/static/wasm_exec.js
else
	cp "$root/misc/wasm/wasm_exec.js" ../web/static/wasm_exec.js
fi
ls -l ../web/static/mms.wasm
