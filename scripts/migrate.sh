#!/bin/bash
set -e
source .env

OPTION=$1
IMAGE=migrate/migrate

case $OPTION in
    up)
    echo "[+] migrate up"
    docker run --rm -v ${PWD}/scripts/migrations:/migrations --network host ${IMAGE} -path=/migrations/ -database "${DATABASE_URL}" -verbose up
    ;;

    down)
    echo "[+] migrate down"
    docker run --rm -v ${PWD}/scripts/migrations:/migrations --network host ${IMAGE} -path=/migrations/ -database "${DATABASE_URL}" -verbose down -all
    ;;

    *)
    echo "[+] invalid option"
esac

# docker rmi -f ${IMAGE}
