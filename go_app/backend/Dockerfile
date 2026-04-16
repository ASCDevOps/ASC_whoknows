# Build Stage
FROM golang:1.26.2-alpine AS builder

WORKDIR /build

# Copy go.mod/go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN go build -o app .


# Runtime Stage
FROM alpine:3.21

WORKDIR /app

# make dir for sqlite
RUN mkdir -p /app/data

RUN apk upgrade --no-cache

# Copy src from build
COPY --from=builder /build/app .

# Copy templates from build, so app can access them at runtime
COPY --from=builder /build/templates ./templates

# Documentation
EXPOSE 8080

# Entrypoint
CMD ["./app"]
