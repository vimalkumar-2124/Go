#!/bin/bash
DOCKER_REG=vimalkumar2124
# VERSION="$(git describe)"
MODIFIED="$([ -n "$(git status --porcelain "$(dirname "${BASH_SOURCE[0]}")")" ] && echo "M")"
TAG="$DOCKER_REG/go_mysql:$VERSION:$MODIFIED"

docker build -f Dockerfile -t "$TAG" .

if [ $? -eq 0 ]
then
    echo "Pushing $TAG"
    docker push "$TAG"
fi