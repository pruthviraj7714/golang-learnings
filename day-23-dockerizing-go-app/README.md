# Day 23 - Dockerizing Go App

## What I Built

- Dockerized Go API
- PostgreSQL container setup
- docker-compose integration

## How to Run

Whenever you make changes to your Go code, remember to run `docker compose up -d --build` to force Docker to build the latest version of your code before starting the container!

```bash
# Start the application
docker compose up -d --build

# View application logs
docker compose logs -f app

# View database logs
docker compose logs -f db

# Stop the application
docker compose down -v
```

## Key Learnings

- Containerization
- Service networking
- Environment variables
