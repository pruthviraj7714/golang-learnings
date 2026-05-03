# Day 22 - Database Migrations

## What I Learned

- Schema versioning
- Migration workflow
- Using golang-migrate

## Key Takeaway

Database schema should be version-controlled like code.

## Commands To follow

i> Start a PostgreSQL container

```
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres
```

ii> Install migrate CLI (for mac if you already have install skip this command)

```
brew install migrate
```

iii> Create migrations files

```
migrate create -ext sql -dir migrations "<migration-name>"
```

iv> Apply migrations

```
migrate -path migrations -database "postgres://postgres:password@localhost:5432" up
```

v> Down migrations

```
migrate -path migrations -database "postgres://postgres:password@localhost:5432" down
```

vi> Version check

```
migrate -path migrations -database "postgres://postgres:password@localhost:5432" version
```

## Practice

- Created users table using migration
- Ran up & down migrations
- Created products and reviews table using migration
- Ran up & down migrations
