FROM golang:1.20 AS builder
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./tictactoe

#second stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/tictactoe .
CMD ["./tictactoe"]