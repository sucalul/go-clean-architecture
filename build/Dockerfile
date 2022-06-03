FROM golang:1.17-alpine as base

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go build cmd/app/main.go
