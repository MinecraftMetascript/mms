#!/usr/bin/env bash
# Checks that keep the grammar honest beyond the corpus tests:
#  - src/ is what `tree-sitter generate` produces from grammar.js;
#  - every example and engine test file (except the deliberately broken
#    recovery inputs) parses without errors;
#  - every Zed query compiles against the grammar;
#  - the Zed copy of highlights.scm matches queries/highlights.scm.
set -euo pipefail
cd "$(dirname "$0")/.."
root=../..
ts="npx --no-install tree-sitter"

$ts generate
if ! git diff --quiet -- src; then
  echo "src/ is out of date: run 'npx tree-sitter generate' and commit the result." >&2
  git --no-pager diff --stat -- src >&2
  exit 1
fi

mapfile -t files < <(find "$root/examples" "$root/engine" -name '*.mms' -not -path '*/recovery/*' | sort)
$ts parse --quiet "${files[@]}"
echo "parsed ${#files[@]} files without errors"

for q in "$root"/editors/zed/languages/mms/*.scm; do
  $ts query "$q" "$root/examples/badlands/badlands.mms" > /dev/null
done
echo "queries compile"

diff -u queries/highlights.scm "$root/editors/zed/languages/mms/highlights.scm"
echo "highlights.scm copies match"
