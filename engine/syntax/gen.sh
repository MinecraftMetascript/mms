#!/bin/sh
# Regenerates the ANTLR parser. Run from engine/syntax. CI runs this and
# fails if the checked-in code differs.
set -e
JAR=.cache/antlr-4.13.1-complete.jar
if [ ! -f "$JAR" ]; then
  mkdir -p .cache
  curl -sSL -o "$JAR" https://www.antlr.org/download/antlr-4.13.1-complete.jar
fi
cd grammar
java -jar ../$JAR -Dlanguage=Go -package parser -o ../parser -Xexact-output-dir -no-listener -no-visitor MMSLexer.g4 MMSParser.g4
