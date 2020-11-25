#!/bin/bash -xe

DOCKER_USER_ID="$(id -u):$(id -g)"

docker_golang() {
  docker volume create desolation_golang_dev --label desolation_dev 1>&2 || true
  docker run $DESOLATION_DOCKER_EXTRA --rm --label desolation_dev \
    --mount source=desolation_golang_dev,target=/go \
    -v "$HOME/.config/gcloud:/go/.config/gcloud" \
    -v "$(pwd)":/go/src -w /go/src \
    -p 8080:8080 -e "HOME=/go" -u "$DOCKER_USER_ID" \
    golang:1.14 "$@"
}

docker_devvnc() {
  docker volume create desolation_vscode_dev --label desolation_dev 1>&2 || true
  docker run $DESOLATION_DOCKER_EXTRA --rm --label desolation_dev \
    --mount source=desolation_vscode_dev,target=/home/hexblade/.vscode \
    -v "$(pwd)":/home/hexblade/desolation \
    -p 5900:5900 \
    murer/desolation-dev:local "$@"
}

docker_devx() {
  docker volume create desolation_vscode_dev --label desolation_dev 1>&2 || true
  docker run $DESOLATION_DOCKER_EXTRA --rm --label desolation_dev \
    --mount source=desolation_vscode_dev,target=/home/hexblade/.vscode \
    -v "$(pwd)":/home/hexblade/desolation \
    -e "DISPLAY=unix$DISPLAY" \
    -v "/tmp/.X11-unix:/tmp/.X11-unix" \
    -p 5010:5010 \
    murer/desolation-dev:local "$@"
}

cmd_code() {
  docker rm -f desolation-vscode || true
  DESOLATION_DOCKER_EXTRA="--name desolation-vscode" cmd_rund devx code --verbose desolation
}

cmd_cleanup() {
  docker ps -aq --filter label=desolation_dev | xargs docker rm -f || true
  docker system prune --volumes --filter label=desolation_dev -f || true
}

cmd_run() {
  dockername="${1?'docker name'}"
  shift
  "docker_${dockername}" "$@"
}

cmd_runi() {
  istty=-i
  [[ -t 0 ]] && istty=-it
  DESOLATION_DOCKER_EXTRA="$DESOLATION_DOCKER_EXTRA $istty" cmd_run "$@"
}

cmd_rund() {
  DESOLATION_DOCKER_EXTRA="$DESOLATION_DOCKER_EXTRA -dit" cmd_run "$@"
}

cmd_devimg() {
  docker build -t murer/desolation-dev:local .
}

cd "$(dirname "$0")"; _cmd="${1?"cmd is required"}"; shift; "cmd_${_cmd}" "$@"
