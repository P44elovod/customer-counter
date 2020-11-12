.SILENT:
run:
	go run ./cmd/server/main.go

build:
	go build -v ./cmd/server


test:
	go test -v -timeout 30s ./...

.DEFAULT_GOAL := build
.PHONY: test