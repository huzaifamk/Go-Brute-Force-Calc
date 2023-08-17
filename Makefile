.PHONY: app

build:
	@echo "Building app..."
	@go mod tidy
	@go build -o bin/main main.go

run:
	@echo "Running app..."
	@./bin/main