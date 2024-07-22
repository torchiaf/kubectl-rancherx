#!/bin/bash

# Helper script to run Rancher in a Docker container

VERSION=$1

echo "Starting container ..."
ID=$(docker run -d --restart=unless-stopped -p 80:80 -p 443:443 --privileged -e CATTLE_BOOTSTRAP_PASSWORD=admin rancher/rancher:${VERSION})

if [ $? -ne 0 ]; then
  echo "An error occurred running the Docker container"
  exit 1
fi

echo ""
echo "Container Id: ${ID}"

echo ""
echo "Waiting for backend to become ready..."
TIME=0
while [[ "$(curl --insecure -s -m 5 -o /dev/null -w ''%{http_code}'' https://localhost)" != "200" ]]; do
  sleep 5;
  TIME=$((TIME + 5))
  printf "\r${TIME}s ... "
done

RANCHER_TOKEN=$(curl -vkL \
  -H 'Content-Type: application/json' \
  -d '{"description":"UI session","responseType":"cookie","username":"admin","password":"admin"}' \
  -X POST \
  https://localhost/v3-public/localProviders/local?action=login 2>&1 | \
  awk -F'[=;]' '/Set-Cookie/{gsub(" ","",$2);print $2;exit}')

echo ""
echo "Get kubeconfig from local cluster ..."

HEADER=$(echo "Cookie: R_SESS=$RANCHER_TOKEN")
curl -kL \
  -H 'Content-Type: application/json' \
  -X POST \
  -H "$HEADER" \
  https://localhost/v3/clusters/local?action=generateKubeconfig | \
  yq '.config' > kubeconfig.yaml

export KUBECONFIG=kubeconfig.yaml

sleep 2;

echo ""
kubectl cluster-info
if [ $? -ne 0 ]; then
  echo "Unable to get Rancher resources"
  exit 1
fi

echo "Done"