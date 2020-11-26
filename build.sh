#!/bin/bash -xe

cmd_main() {
  go run main.go "$@"
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
