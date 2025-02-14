build:
	@go build -o bin/api .

run: build
	@./bin/api

test:
	go test -v ./redis_helpers_test.go redis_helpers.go redis_datastore.go datastore.go data.go

redis:
	docker run --name redis -d -p  5000:6379 redis:latest
