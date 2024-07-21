#!/bin/bash

# Helper script to run Rancher in a Docker container

VERSION=$1

echo "Starting container ..."
ID=$(docker run -d --restart=unless-stopped -p 80:80 -p 443:443 --privileged -e CATTLE_BOOTSTRAP_PASSWORD=admin rancher/rancher:${VERSION})

echo "Container Id: ${ID}"
