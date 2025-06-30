#!/bin/sh
set -e

echo "Running migrations..."
goose -dir ./app/database/migrations postgres "${DB_DSN}" up

echo "Starting application..."
exec ./app/main
