#!/bin/sh

docker rm -f -v $(docker ps -a -q)