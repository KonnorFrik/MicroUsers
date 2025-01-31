#! /usr/bin/bash
set -e

docker compose down
sudo docker compose build
docker compose up -d
