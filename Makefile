.PHONY: build run clean test fmt vet help

BINARY_NAME=playground

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go

clean:
	go clean
	rm -f $(BINARY_NAME)

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

deps:
	go mod download
	go mod tidy

build-run: build
	./$(BINARY_NAME)

help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  clean      - Remove build artifacts"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format code"
	@echo "  vet        - Run go vet"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  build-run  - Build and run the application"
	@echo "  help       - Show this help message"
