docker run --name redis-test -p 6379:6379 -d redis
go run cmd/main.go