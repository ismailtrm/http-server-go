#!/bin/sh
#
# Build and run the HTTP server
#

set -e # Exit early if any commands fail

echo "Building HTTP server..."
go build -o /tmp/http-server-go app/*.go

echo "Starting server on 0.0.0.0:4221..."
exec /tmp/http-server-go "$@"
