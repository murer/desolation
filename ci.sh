#!/bin/bash -xe

cmd_script() {
  ./build.sh test ./...
  ./build.sh build_all
}

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
