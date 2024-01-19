#!/bin/ash
# Specifies a parent image
FROM golang:1.21-alpine AS builder

# Creates an app directory to hold your app’s source code
WORKDIR /build

# Installs Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Builds your app with optional configuration
RUN GOARCH=amd64 GOOS=linux go build -o app ./app.go

FROM scratch

WORKDIR /test/

COPY --from=builder /build/app .
COPY --from=builder /build/config.yml .

# Tells Docker which network port your container listens on
EXPOSE 8082
#ENTRYPOINT ["/app/app"]
CMD ["/test/app"]


