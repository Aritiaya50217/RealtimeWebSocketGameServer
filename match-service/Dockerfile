FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o match-service ./cmd/match/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/match-service .
CMD ["./match-service"]
