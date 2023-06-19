all: fmt lint test

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golint ./...
