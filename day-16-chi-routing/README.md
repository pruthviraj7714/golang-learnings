# Day 16 - Routing with Chi

## What I Built

- API using chi router
- Dynamic routes (/users/{id})
- Route grouping (/users)

## Key Learnings

- Clean routing in Go
- URL params handling
- Struct-based JSON response
- route grouping

## Endpoints

- GET / → Hello from chi router
- GET /hello?username=John → Hello, John
- GET /user → JSON response
- GET /users → Users list
- GET /api/users/{id} → get user by id
