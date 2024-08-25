FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o cli

FROM alpine:latest

COPY --from=builder /app/cli /usr/local/bin/cli

ENTRYPOINT ["/usr/local/bin/cli"]
