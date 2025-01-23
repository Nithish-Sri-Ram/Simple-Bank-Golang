#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
/app/migrate -path ./migration -database "$DB_SOURCE" up

echo "start the app"
exec "$@"
# the above means takes all the parameters passed to the script and run it 