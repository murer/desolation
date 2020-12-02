#!/bin/bash -xe

cmd_main() {
  go run main.go "$@"
}

_static_gen() {
  set +x
  echo "package public"
  echo "func init() {"
  ls guest/public | grep -v "\.go$" | while read k; do
    echo "StaticFiles[\"$k\"] = \`$(cat "guest/public/$k")\`"
  done
  echo "}"
  set -x
}

cmd_static_gen() {
  _static_gen > guest/public/gen_public.go
}

cmd_build() {
  cmd_static_gen
  desolation_goos="${1?'use: linux, darwin or windows'}"
  desolation_goarch="${2:-"amd64"}"
  desolation_ext=""
  if [[ "x$desolation_goos" == "xwindows" ]]; then desolation_ext=".exe"; fi
  desolation_ldflags="-s -w -extldflags '-static'"
  mkdir -p build
  CGO_ENABLED="0" GOOS="$desolation_goos" GOARCH="$desolation_goarch" \
    go build -a -trimpath -ldflags "$desolation_ldflags" \
      -installsuffix cgo -tags netgo -mod mod \
      -o "build/desolation-$desolation_goos-${desolation_goarch}${desolation_ext}" .
  du -hs "build/desolation-$desolation_goos-${desolation_goarch}${desolation_ext}"
}

cmd_build_all() {
  rm -rf build
  cmd_build linux
  cmd_build darwin
  cmd_build windows
  _build_enc
  cd build
  sha256sum -b * > SHA256
  cd -
}

_build_enc() {
  ls build | while read k; do 
    cat "build/$k" | base64 > "build/$k.base64.txt"
  done
}

cmd_test() {
  go clean -testcache
  go test "$@"
}

cmd_fmt() {
   go fmt "$@"
}

cmd_vendor() {
  go mod vendor -v
}

cmd_test_ssh() {
  ssh -o "ProxyCommand go run main.go guest %h %p" "$@"
}

cmd_cap() {
  sleep 5 && (import -window root png:- | zbarimg png:-)
}

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
