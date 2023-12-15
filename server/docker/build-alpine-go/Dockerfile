FROM golang:1.21-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify