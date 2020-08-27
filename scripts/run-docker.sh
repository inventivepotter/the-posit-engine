#!/bin/bash
set -x

docker images -a | grep "the-posit-engine:latest" | awk '{print $3}' | xargs docker rmi
docker build . -t the-posit-engine:latest
docker run -p 1323:1323 --rm the-posit-engine:latest