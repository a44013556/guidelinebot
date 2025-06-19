FROM golang:1.21-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

# 執行
CMD ["./main"]
