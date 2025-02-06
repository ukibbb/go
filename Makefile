build:
	@go build -o bin/api .

run: build
	@./bin/api

test:
	go test -v ./...

redis:
	docker run --name redis -d -p  5000:6379 redis:latest
