#!/bin/bash

PROJECT_DIR=$(dirname "$0")
GENERATE_DIR="$PROJECT_DIR/cmd/generate"

echo "Start Generating"
go run $GENERATE_DIR