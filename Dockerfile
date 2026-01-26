# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

COPY --from=builder /app/database/migrations ./database/migrations
# optional: install CA certs for HTTPS if using Render
RUN apk add --no-cache ca-certificates

ENV PORT=8080
CMD ["./main"]
