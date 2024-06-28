build:
	@go build -o bin/api

run: build
	@./bin/api --listenAddr :8000

test:
	@go test -v ./...