FROM golang:1.21-bullseye AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine/git:latest
FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x /app/main

CMD ["./main"]
