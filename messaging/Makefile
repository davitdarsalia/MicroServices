.PHONY: all clean race vet test testVerbose format build run

BINARY_NAME=auth

all: clean race vet test build

clean:
	go clean
	if [ -f $(BINARY_NAME) ]; then rm -f $(BINARY_NAME); fi

race:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

vet:
	go vet ./...

#test:
#	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

#testVerbose:
#	go test -v -race -coverprofile=coverage.txt -covermode=atomic -memprofile mem.out ./...

format:
	go fmt ./...

build:
	cd cmd && go build -o $(BINARY_NAME) cmd/main.go

run:
	cd cmd && ./$(BINARY_NAME)

lint:
	golangci-lint run --config .golangci.yaml ./...
