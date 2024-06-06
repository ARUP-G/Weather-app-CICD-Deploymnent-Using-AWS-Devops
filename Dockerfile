# Base image
FROM golang:1.21-alpine as builder

# Working directory
WORKDIR /Weather-app

# Copy go mod and sum file
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy source code from current dir to working dir
COPY . .

# Build go app
RUN go build -o main .

# 2nd Stage
FROM alpine:latest

# Set current working dir n the container
WORKDIR /root/

# Copy pre-built binary file freom the base stage
COPY --from=builder /Weather-app/main .
COPY --from=builder /Weather-app/static ./static
COPY --from=builder /Weather-app/.env .

# Expose port to external world
EXPOSE 8080

# Command to run the executable
CMD [ "./main" ]