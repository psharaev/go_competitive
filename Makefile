.PHONY: build
build:
	go build -o build/goc ./cmd/goc

test:
	go test ./...
