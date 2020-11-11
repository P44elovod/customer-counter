.SILENT:
run:
	go run ./cmd/server/main.go

build:
	go build -v ./cmd/server


.PHONY: test
test:
	go test -v -rase -timeout 30s ./...

.DEFAULT_GOAL := build