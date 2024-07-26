#!/usr/bin/env bash
mkdir -p build
docker run --network=host --rm --mount type=bind,source=.,target=/workdir -w /workdir golang:alpine go build -o build/uttt
# docker rmi golang:alpine # Uncomment if you do not want to keep the image on you machine
