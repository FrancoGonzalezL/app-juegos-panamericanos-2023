


.PHONY: build run
build: 
	@echo "Building..."
	@go build -o bin/main cmd/main.go

run: build
	@./bin/main


