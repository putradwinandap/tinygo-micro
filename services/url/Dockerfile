# Use official Go image as builder
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go mod files and download deps first (build cache optimization)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o url-service ./cmd/main.go

# Final image
FROM alpine:latest

WORKDIR /app

# Copy binary only
COPY --from=builder /app/url-service .

# Port (ubah jika perlu)
EXPOSE 8002

CMD ["./url-service"]
