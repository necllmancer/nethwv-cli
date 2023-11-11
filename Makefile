# Makefile for building the nethwv CLI tool

# Golang commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# binary
BINARY_NAME=nethwv

# build directory
BUILD_DIR=./build

.PHONY: all build clean install run test deps

all: test build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/nethwv/main.go

test: 
	$(GOTEST) -v ./...

clean: 
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

deps:
	$(GOGET) ./...

install: build
	@echo "Installing..."
	mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

run: build
	@echo "Running..."
	@./$(BUILD_DIR)/$(BINARY_NAME)
