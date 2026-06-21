# FROM golang:1.25-alpine AS builder
# RUN apk update && apk add --no-cache git
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o ./bin/main ./cmd/api/main.go
# CMD ["./bin/main"]

FROM golang:1.25-alpine

RUN apk add --no-cache git

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air"]
