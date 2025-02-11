FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bot ./cmd/

WORKDIR /app

CMD ["./bot"]
