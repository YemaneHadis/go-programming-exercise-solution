#!/bin/bash

# Check if a filename was provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <filename.go>"
  exit 1
fi

# Get the filename from the argument
FILE_NAME=$1

# Check if the file exists
if [ ! -f "$FILE_NAME" ]; then
  echo "File $FILE_NAME does not exist."
  exit 1
fi

# Create the build directory if it doesn't exist
BUILD_DIR="build"
if [ ! -d "$BUILD_DIR" ]; then
  mkdir $BUILD_DIR
fi

# Build the Go file
go build -o $BUILD_DIR/$(basename "$FILE_NAME" .go) $FILE_NAME

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful! The binary is located in the $BUILD_DIR directory."
else
  echo "Build failed."
fi
