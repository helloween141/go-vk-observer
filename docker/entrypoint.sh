#!/bin/sh
set -e

echo "Waiting for PostgreSQL..."
timeout 60 sh -c "until pg_isready -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USERNAME}; do sleep 2; done"

DSN="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=${DB_SSL_MODE}"
echo "Running migrations..."
goose -dir ./app/database/migrations postgres "${DSN}" up

echo "Starting application..."
exec /main