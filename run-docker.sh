#!/bin/bash
set -e

debug=$1

docker-compose down --remove-orphans
docker-compose rm -fv

rm -rf ./postgres-data

docker volume create redpanda1 && \
docker volume create redpanda2 && \
docker volume create redpanda3


if [ -z "$debug" ]
then
  docker-compose up --build web
else
  docker-compose up --build web-debug
fi
