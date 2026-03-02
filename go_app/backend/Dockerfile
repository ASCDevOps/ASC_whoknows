# Build Stage
FROM golang:1.25.7-alpine AS builder

WORKDIR /build

# Copy go.mod/go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN GOOS=linux GOARCH=amd64 go build -o app


# Runtime Stage
FROM alpine:3.21

WORKDIR /

# Copy src from build
COPY --from=builder /build/app .
# Copy templates from build, so app can access them at runtime
COPY --from=builder /build/templates ./templates

# Documentation
EXPOSE 8080

# Entrypoint
CMD ["./app"]