#!/bin/bash
# Run database migrations

# Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
    echo "Error: DATABASE_URL environment variable is not set"
    exit 1
fi

# Default to "up" if no argument provided
DIRECTION=${1:-up}

echo "Running migrations $DIRECTION..."
migrate -path ./migrations -database "$DATABASE_URL" $DIRECTION
echo "Done!"
