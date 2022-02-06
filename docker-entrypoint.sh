#!/bin/ash
set -eo pipefail

echo "Creating config file from env vars"

if test -z "${SERVER_JWT_SECRET}"; then
  SERVER_JWT_SECRET="$(head -n 10 /dev/urandom | md5sum | cut -c 1-32)"
  export SERVER_JWT_SECRET
fi

if test -z "${SERVER_REDIS_ADDRESS}"; then
  SERVER_REDIS_ADDRESS="redis:6379"
  export SERVER_REDIS_ADDRESS
fi

if test -z "${SERVER_REDIS_PASSWORD}"; then
  SERVER_REDIS_PASSWORD=""
  export SERVER_REDIS_PASSWORD
fi

if test -z "${SERVER_DB_ADDRESS}"; then
  SERVER_DB_ADDRESS="postgres"
  export SERVER_DB_ADDRESS
fi

if test -z "${SERVER_DB_USER}"; then
  SERVER_DB_USER="postgres"
  export SERVER_DB_USER
fi

if test -z "${SERVER_DB_PASSWORD}"; then
  SERVER_DB_PASSWORD="0f359740bd1cda994f8b55330c86d845"
  export SERVER_DB_PASSWORD
fi

if test -z "${SERVER_DB_NAME}"; then
  SERVER_DB_NAME="tinyblog"
  export SERVER_DB_NAME
fi

rm -f /tmp/tmp.yaml
(
  echo "cat <<EOF >/app/config.yaml"
  cat /app/config.docker.yaml
  echo "EOF"
) >/tmp/tmp.yaml

. /tmp/tmp.yaml

if [ ! -f /app/inited ]; then
  echo "Running install"
  /app/server migrate
  touch /app/inited
fi

exec /app/server "$@"
