# Сборка
FROM golang:1.26-rc-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server main.go

# Финальный легкий образ
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
