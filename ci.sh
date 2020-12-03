#!/bin/bash -xe

cmd_script() {
  ./docker.sh build
}

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
