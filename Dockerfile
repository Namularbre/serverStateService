# Step 1: Build the Go application
# Use the official Golang image to build the application
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files (if present)
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod tidy

# Copy all source code into the working directory
COPY . .

# Build the application binary
RUN GOOS=linux GOARCH=amd64 go build -o myapp

# Expose the port where the application will listen (e.g., 8080)
EXPOSE 8080

# Set the environment variable for the MariaDB connection string
ENV CONN_STR="mariadb://namu:2261@localhost:3306/server_monitoring"
# Set the environment variable for the http port
ENV PORT=8080
# Set the environment variable for the http address
ENV ADDRESS=localhost

# Default command to run the application
CMD ["./myapp"]
