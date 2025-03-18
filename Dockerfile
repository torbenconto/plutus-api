# First stage: build the application
FROM golang:1.22-bookworm AS builder

WORKDIR /app

LABEL org.opencontainers.image.source=https://github.com/torbenconto/plutus-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Second stage: copy the binary into a smaller base image
FROM debian:bookworm-slim

WORKDIR /app

# Install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Create an unprivileged user
RUN useradd -u 10001 appuser

# Copy the binary from the first stage
COPY --from=builder /app/main /app/main

# Use the unprivileged user to run the application
USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/main"]