dev:
	@docker run --name mongodb -d -p 27017:27017 mongo:latest

build:
	@go build -o bin/api

run: build
	@./bin/api --listenAddr :8000

test:
	@go test -v ./...