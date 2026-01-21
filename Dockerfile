# Build Stage
ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} AS builder

WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY pkg ./pkg

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o deb-tester ./cmd/deb-tester

# Test Stage (Runtime)
FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

# Install essential tools
RUN apt-get update && apt-get install -y \
    coreutils \
    curl \
    git \
    dpkg \
    apt-utils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/deb-tester /usr/local/bin/deb-tester

ENTRYPOINT ["deb-tester"]
