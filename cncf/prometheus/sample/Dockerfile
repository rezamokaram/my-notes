# Use the official Golang image to create a build artifact
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN go build -o http-server
ENTRYPOINT ["/app/http-server"]

# Use a minimal image to run the application
# FROM alpine:latest

# Copy the binary from the builder stage
# COPY --from=builder /app/http-server /usr/local/bin/http-server

# Expose the application port
# EXPOSE 8080

# Command to run the application
# ENTRYPOINT ["/usr/local/bin/http-server"]
