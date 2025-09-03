# Get the short git commit hash (7 characters)
COMMIT_HASH := $(shell git rev-parse --short=7 HEAD)

BUILD_DIR := build

# Build the Go binary with the commit hash embedded
build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.commitHash=$(COMMIT_HASH)" -o $(BUILD_DIR)/shortUrls

# Test the application and generate HTML coverage file
test:
	mkdir -p $(BUILD_DIR)
	go test ./... -cover -coverprofile=$(BUILD_DIR)/coverage.out
	go tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html

# Run the application
run: build
	./build/shortUrls

# Clean up generated files
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build test run clean