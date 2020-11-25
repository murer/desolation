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

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
