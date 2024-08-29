FROM golang:1.22.0-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o build/plasma-da ./cmd/plasma-da

# Create a minimal image
FROM alpine:3.12.0 AS runner

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/build/plasma-da /app/build/plasma-da

# Run
CMD ["build/plasma-da", "start"]