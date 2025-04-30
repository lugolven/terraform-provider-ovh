# Pinning the version of various tools
FROM golang:1.24 AS go
FROM golangci/golangci-lint:v2.1.5 AS golangci

# Base image for building the provider
FROM ubuntu:25.04 AS base
RUN apt update && apt install make git -y
COPY --from=go /usr/local/go/ /usr/local/go/
COPY --from=golangci /usr/bin/golangci-lint /usr/bin/golangci-lint
ENV PATH="/usr/local/go/bin:${PATH}"

# Image to develop the provider
FROM base AS dev
RUN apt update && apt install -y bash-completion curl
RUN useradd -m -s /bin/bash dev
USER dev

# Image to run during continuous integration
FROM base AS ci
RUN useradd -m -s /bin/bash ci && \
    mkdir -p /build && \
    chown -R ci:ci /build

WORKDIR /build
COPY . .
USER ci