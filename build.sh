#!/usr/bin/env bash
die() { echo "[FATAL] $1"; exit 1; }
for dep in "git" "go" "zip"; do
  hash "$dep" 2>/dev/null || die "missing $dep :("
done

commit=$(git rev-parse HEAD 2>/dev/null) || die 'unable to get commit hash'
oss=("darwin" "linux" "windows")
arch="amd64" # no 386 comps in these woods folks
name="emdee"

echo "--- building for [${oss[*]}] [$arch]"
for os in "${oss[@]}"; do
  echo "~~~ building $name-$os"
  [[ $os == "windows" ]] && bin="$name.exe" || bin="$name"
  GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 \
    go build -ldflags "-X main.commit=$commit" -o "$bin" . || die 'cant build'
  zip "$name-$os.zip" "$bin" || die "zip problems: $name-$os"
done
