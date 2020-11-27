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

cmd_build() {
  _static_gen > guest/public/gen_staticfiles.go
  mkdir -p target
  go build -o target/desolation
}

cmd_test() {
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

cmd_resource2go() {
  ls guest/public | while read k; do
    echo "package public"
  done
}

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
