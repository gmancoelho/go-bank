# The binaries directory
BIN_DIR=bin

# The executable name
EXECUTABLE=go-bank

# A phony target is one that is not really the name of a file
.PHONY: all build run test clean

# The default target
all: build

# Build the main Go binary
build:
	@echo "Building the Go application..."
	@go build -o ./$(BIN_DIR)/$(EXECUTABLE) cmd/main.go

# Run the built Go binary
run: build
	@echo "Running the Go application..."
	@./$(BIN_DIR)/$(EXECUTABLE)

# Run the tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean up the build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)/$(EXECUTABLE)