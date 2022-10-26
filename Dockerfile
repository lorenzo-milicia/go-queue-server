FROM golang:1.18-alpine

WORKDIR /src

COPY . ./

RUN go build -o /queue-server
