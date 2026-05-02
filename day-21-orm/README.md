# Day 21 - ORM vs Raw SQL

## What I Learned

- GORM (ORM in Go)
- SQLC (type-safe queries)
- Tradeoffs between approaches

## Key Takeaway

SQLC provides best balance of performance and safety for production systems.

## Basic CRUD Operations Using GORM

```go
//get all users
  db.Find(&users)

//get user by id
  db.First(&user, "id = ?", userId)

//create user
  db.Create(&user)

//delete user
  db.Delete(&user)

//update user
  db.Model(&user).Where("id = ?", userId).Update(&updatedUser)
```

## Command To Install Package

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

## Practice

- Converted CRUD API to GORM
