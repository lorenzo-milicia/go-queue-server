FROM golang:1.18-alpine

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download