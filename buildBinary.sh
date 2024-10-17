#!/bin/bash

# Set the version
VERSION="v0.1"

# Create binary folder if it doesn't exist
mkdir -p binary

echo "Building Nox for Windows..."
GOOS=windows GOARCH=amd64 go build -o binary/nox_${VERSION}_windows_amd64.exe

echo "Building Nox for Mac (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o binary/nox_${VERSION}_mac_amd64

echo "Building Nox for Mac (M1)..."
GOOS=darwin GOARCH=arm64 go build -o binary/nox

echo "Build complete! Binaries stored in the 'binary' folder."
