FROM golang:1.20-alpine3.16 AS builder

WORKDIR /app
COPY . .

RUN GOOS=linux go build -o ./bin/app ./cmd/main.go

# Run stage
FROM alpine:latest AS runner

COPY --from=builder /app/bin/app/ .
COPY --from=builder /app/.env .

EXPOSE 9000
CMD ["./app"]