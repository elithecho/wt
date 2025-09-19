.PHONY: build clean install uninstall test help

# Binary name
BINARY_NAME=wt
BUILD_DIR=build
INSTALL_PATH=/usr/local/bin

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "✅ Built $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Cleaned"

# Install globally
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_PATH)/
	@echo "✅ Installed! You can now use 'wt' from anywhere"

# Uninstall
uninstall:
	@echo "Removing $(BINARY_NAME) from $(INSTALL_PATH)..."
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "✅ Uninstalled"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Show help
help:
	@echo "Available commands:"
	@echo "  build     - Build the binary to $(BUILD_DIR)/$(BINARY_NAME)"
	@echo "  clean     - Remove build artifacts"
	@echo "  install   - Build and install globally to $(INSTALL_PATH)"
	@echo "  uninstall - Remove from $(INSTALL_PATH)"
	@echo "  test      - Run tests"
	@echo "  help      - Show this help"

# Default target
all: build