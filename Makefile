.DEFAULT_GOAL := build
.PHONY: fmt vet build test clean tidy run-server run-client run-server-debug run-client-debug

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./...

test: build
	go test ./...

clean:
	go clean ./...

tidy:
	go mod tidy

run-server:
	go run ./examples/server

run-client:
	go run ./examples/client

run-server-debug:
	dlv debug ./examples/server --

run-client-debug:
	dlv debug ./examples/client -- -debug
