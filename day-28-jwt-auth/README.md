# Day 28 - JWT Authentication

## What I Built

- Signup/login system
- JWT authentication
- Protected routes
- RBAC middleware

## Commands to install

```bash
# Install jwt
go get github.com/golang-jwt/jwt/v5

# Install chi router
go get github.com/go-chi/chi/v5

# Install bcrypt
go get golang.org/x/crypto/bcrypt
```

## Key Learnings

- Password hashing with bcrypt
- JWT token generation
- Auth middleware
- Role-based access control

## Common Backend Flow

### 1. Signup → Hash Password → Store User

### 2. Login → Verify Password → Issue JWT

### 3. Protected API → Validate JWT

## Endpoints

- POST /signup
- POST /login
- GET /profile
