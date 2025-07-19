#!/bin/bash

# Tedo Installation Script
echo "Installing Tedo..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

# Install Tedo
echo "Installing Tedo from GitHub..."
go install github.com/WasathTheekshana/tedo/cmd/tedo@latest

# Check if installation was successful
if command -v tedo &> /dev/null; then
    echo "✅ Tedo installed successfully!"
    echo "Run 'tedo' to start the application."
else
    echo "❌ Installation failed. Make sure \$GOPATH/bin is in your \$PATH"
    echo "Add this to your shell profile:"
    echo "export PATH=\$PATH:\$(go env GOPATH)/bin"
fi
