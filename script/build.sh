#!/usr/bin/env bash
set -eu
for dep in "go" "git"; do
  hash "$dep" 2>/dev/null || { echo "missing $dep :("; exit 1; }
done
commit=$(git rev-parse HEAD 2>/dev/null)
go build -ldflags "-X main.commit=$commit" -o emdee .
