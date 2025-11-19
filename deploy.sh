#!/bin/bash

echo "🚀 Starting deployment..."

# Pull latest changes
echo "📥 Pulling latest code..."
git pull

# Check for docker-compose or docker compose
if command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
elif docker compose version &> /dev/null; then
    DOCKER_COMPOSE="docker compose"
else
    echo "❌ Error: docker-compose or docker compose plugin not found."
    exit 1
fi

# Build and start containers
echo "🐳 Building and starting containers using $DOCKER_COMPOSE..."
$DOCKER_COMPOSE up -d --build

if [ $? -eq 0 ]; then
    echo "✅ Deployment complete! App is running on port 3000."
else
    echo "❌ Deployment failed."
    exit 1
fi
