FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o service-account ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service-account .
COPY .env .
CMD ["./service-account"]