# Build Stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY pkg ./pkg

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o deb-tester ./cmd/deb-tester

# Test Stage (Runtime)
FROM debian:stable-slim

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/deb-tester /usr/local/bin/deb-tester

ENTRYPOINT ["deb-tester"]
