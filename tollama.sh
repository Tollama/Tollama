#!/bin/bash

# Install Ollama
sudo pacman -S ollama

# Start the Ollama server in the background
ollama serve &

# Wait for a few seconds to ensure the server is up and running
sleep 5

# Verify the server is running
curl http://localhost:11434/api/tags

# Run the Go application
go run src/*.go