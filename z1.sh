#!/bin/bash

# last is 24
IDX=$1
TAG="172.10.60.2/wolcen/bookstack:1.4-Z$IDX"
docker build -t $TAG .
docker push  $TAG

let LIDX=IDX-1
LTAG="172.10.60.2/wolcen/bookstack:1.4-Z$LIDX"
docker rmi $LTAG
