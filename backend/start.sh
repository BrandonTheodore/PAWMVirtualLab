#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

echo "Running go mod tidy..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "Error running go mod tidy"
    exit 1
fi

echo "Building the application..."
go build -o pawm-virtual-lab
if [ $? -ne 0 ]; then
    echo "Error building the application"
    exit 1
fi

echo "Starting the application..."
./pawm-virtual-lab