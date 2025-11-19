#!/bin/bash

echo "🚀 Starting deployment..."

# Pull latest changes
echo "📥 Pulling latest code..."
git pull

# Build and start containers
echo "🐳 Building and starting containers..."
docker-compose up -d --build

echo "✅ Deployment complete! App is running on port 3000."
