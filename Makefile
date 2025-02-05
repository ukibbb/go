build:
	@go build -o bin/api .

run: build
	@./bin/api

test:
	go test -v ./...

redis:
	docker run --name redis -p -d 5000:6379 redis:latest
