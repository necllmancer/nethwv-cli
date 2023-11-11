# Makefile for building the nethwv CLI tool

# Binary name
BINARY_NAME=nethwv

# Build directory
BUILD_DIR=./build

.PHONY: all build clean install run

all: build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/nethwv/main.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

install: build
	@echo "Installing..."
	mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

run: build
	@echo "Running..."
	@./$(BUILD_DIR)/$(BINARY_NAME)