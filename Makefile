.PHONY: build run test lint

build:
	go build -o bin/core .

run:
	go run .

test:
	go test ./... -v

lint:
	golangci-lint run ./...
