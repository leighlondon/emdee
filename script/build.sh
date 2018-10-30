#!/usr/bin/env bash
set -eu
die_() { echo "[FATAL] $1"; exit 1; }
for dep in "go" "git"; do
  hash "$dep" 2>/dev/null || die_ "missing $dep :("
done

commit=$(git rev-parse HEAD 2>/dev/null) || die_ 'unable to get commit hash'
oss=("darwin" "linux" "windows")
arch="amd64" # no 386 comps in these woods folks

echo "--- building for [${oss[@]}] [$arch]"
for os in "${oss[@]}"; do
  GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 \
    go build -ldflags "-X main.commit=$commit" -o "emdee-$os-$arch" .
done

echo '--- complete'
