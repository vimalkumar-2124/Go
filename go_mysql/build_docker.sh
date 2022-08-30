#!/bin/bash
DOCKER_REG=vimal2124
# VERSION="$(git describe)"
MODIFIED=latest
TAG="$DOCKER_REG/go_mysql:$MODIFIED"

docker build -f Dockerfile -t "$TAG" .

if [ $? -eq 0 ]
then
    echo "Pushing $TAG"
    docker push "$TAG"
fi