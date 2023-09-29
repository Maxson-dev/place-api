#!/bin/bash
# wait-for-postgres.sh

set -e

host="postgres"
shift

sleep 5
until psql -h "postgres" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 10
done

>&2 echo "Postgres is up - executing command"

exec "$@"