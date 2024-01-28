all: build

build:
	@echo "Building..."
	@go build -o bin/icache cmd/api/main.go

run:
	@go run cmd/api/main.go

clean:
	@echo "Cleaning..."
	@rm -f main

watch:
	@air

.PHONY: all build run test clean
