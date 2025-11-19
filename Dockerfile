# Stage 1: Build Frontend
FROM node:18-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm install
COPY frontend/ .
# Ensure API URL is relative for production (or set via ENV if needed, but relative is best for same-origin)
# We might need to adjust api.js to use relative path if serving from same origin
RUN npm run build

# Stage 2: Build Backend
FROM golang:1.23-alpine AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN go build -o server cmd/server/main.go

# Stage 3: Final Image
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates

# Copy backend binary
COPY --from=backend-builder /app/backend/server .
COPY --from=backend-builder /app/backend/.env .

# Copy frontend build to dist directory
COPY --from=frontend-builder /app/frontend/dist ./dist

# Expose port
EXPOSE 3000

# Run the server
CMD ["./server"]
