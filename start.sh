#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
echo "DB_SOURCE=$DB_SOURCE"
/app/migrate -path ./migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
# the above means takes all the parameters passed to the script and run it 