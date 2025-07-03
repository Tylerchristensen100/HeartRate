#!/bin/bash


# Vibe coded with Gemini 2.5 Flash


PACKAGE_NAME="heart_rate"
SOURCE_DIR="." # Or your actual source directory, e.g., "./cmd/heart_rate"
BUILD_DIR="./build" # Define the build directory

# Create the build directory if it doesn't exist
mkdir -p "${BUILD_DIR}"

echo "Building for Windows (AMD64)..."
GOOS=windows GOARCH=amd64 go build -o "${BUILD_DIR}/${PACKAGE_NAME}_windows_amd64.exe" "${SOURCE_DIR}"

echo "Building for macOS (Apple Silicon - ARM64)..."
GOOS=darwin GOARCH=arm64 go build -o "${BUILD_DIR}/${PACKAGE_NAME}_macos_arm64" "${SOURCE_DIR}"

echo "Building for Linux (AMD64)..."
GOOS=linux GOARCH=amd64 go build -o "${BUILD_DIR}/${PACKAGE_NAME}_linux_amd64" "${SOURCE_DIR}"

echo "Cross-compilation complete! Binaries are in the '${BUILD_DIR}' folder."