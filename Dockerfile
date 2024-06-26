FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["/app/main"]
