FROM golang:1.18-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o ./bin/queue-server .

CMD ["./bin/queue-server"]