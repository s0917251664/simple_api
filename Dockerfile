FROM        golang:alpine AS builder
RUN         mkdir -p /api
WORKDIR     /api
COPY        . .
RUN         go mod download
RUN         go build -o api
ENTRYPOINT  ["./api"]