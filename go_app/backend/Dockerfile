# Build Stage
FROM golang:1.25.7-alpine AS builder

RUN apk add --no-cache gcc musl-dev git

WORKDIR /build

# Copy go.mod/go.sum and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app


# Runtime Stage
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /

# Copy binary + templates + handlers
COPY --from=builder /build/app .
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/handlers ./handlers

# Documentation
EXPOSE 8080

CMD ["./app"]