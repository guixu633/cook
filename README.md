# Daily Meal Tracker

A full-stack application to track your daily meals, built with Golang (Backend) and React (Frontend).

## Features
- List daily meals with images.
- Add new meals.
- Delete meals.
- Images stored in Aliyun OSS.
- SQLite database.

## Project Structure
- `backend/`: Golang application (Gin + GORM).
- `frontend/`: React application (Vite + Tailwind CSS).
- `Dockerfile`: Multistage build for production.
- `docker-compose.yml`: Orchestration for deployment.

## Local Development

### Backend
```bash
cd backend
go run cmd/server/main.go
```

### Frontend
```bash
cd frontend
npm run dev
```

## Deployment

### Prerequisites
- Docker and Docker Compose installed on the server.
- `.env` file in `backend/` (or environment variables set in docker-compose).

### One-Click Deploy
1. Clone the repository to your server.
2. Ensure your `backend/.env` file is populated with OSS credentials.
3. Run the deployment script:
   ```bash
   ./deploy.sh
   ```

The application will be available at `http://your-server-ip:3000`.
