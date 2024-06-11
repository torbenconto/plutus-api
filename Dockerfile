# Use the official Go image as the base image
FROM golang:1.22-bookworm

# Set the working directory inside the container
WORKDIR /app

LABEL org.opencontainers.image.source=https://github.com/torbenconto/plutus-api

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8001
EXPOSE 8001

# Set the entry point command to run the built binary
CMD ["./main"]
