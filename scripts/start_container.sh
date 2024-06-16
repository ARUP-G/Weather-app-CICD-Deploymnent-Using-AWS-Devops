#!/bin/bash
set -e

# Pull the Docker image from Docker Hub
docker push ard3dk/weather-app:1.1

# Run the Docker image as a container
docker run -d -p 8080:8080 ard3dk/weather-app:1.1