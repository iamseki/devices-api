FROM golang:1.24 as builder

WORKDIR /app

# Its ok to Copy everything for now, to download dependencies first
COPY . .
RUN go mod download

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devices-api .

# ---- Stage 2: Run ----
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --chmod=755 --from=builder /app/devices-api /app/

# Expose the application port (if applicable)
EXPOSE 8081

# Command to run the application
CMD ["./devices-api"]