.PHONY: build clean tool lint help

all: build

build:
	go build -o ./bin/ ./cmd/server
	go build -o ./bin/ ./cmd/client

lint:
	golint ./...

clean:
	rm -rf ./bin/*
	go clean -i .
