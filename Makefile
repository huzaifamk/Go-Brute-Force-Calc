.PHONY: app

build:
	@echo "Building app..."
	@go mod tidy
	@go build -o bin/app main.go

run:
	@echo "Running app..."
	@./bin/app