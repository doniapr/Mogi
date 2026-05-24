.PHONY: help build run test clean install-deps

help:
	@echo "Available commands:"
	@echo "  make install-deps  - Install dependencies"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make dev          - Run in development mode with hot reload"

install-deps:
	go mod download
	go mod tidy

build:
	go build -o bin/mogi main.go

run:
	go run main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

fmt:
	go fmt ./...

lint:
	golangci-lint run

dev:
	air

docker-build:
	docker build -t mocker:latest .

docker-run:
	docker run -p 8000:8000 --env-file .env mocker:latest

