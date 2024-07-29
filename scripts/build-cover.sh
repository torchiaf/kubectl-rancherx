#!/bin/sh

COVER_DIR="cover"

if [ -n "$1" ]; then
  COVER_DIR=$1
fi

rm -rf $COVER_DIR; mkdir $COVER_DIR

go build -o dist/kubectl-rancherx -cover