FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

# 執行
CMD ["./main"]
