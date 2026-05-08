# Day 27 - Redis Caching

## What I Built

- Redis integration with Go
- API response caching
- Cache invalidation

## Key Learnings

- Redis basics
- Cache-aside pattern
- Performance optimization

## Redis Commands Used

- Set : used to set the key in redis
- Get : used to get the value of the key from redis
- Del : used to delete the key from redis
- Exists : used to check if the key is exists in redis
- Expire : used to set the expiry duration of the key from redis
- Ping : used to check the connection with redis server

## Command to get redis

```base
go get github.com/redis/go-redis/v9
```

## Improvements

- Faster API responses
- Reduced DB load
