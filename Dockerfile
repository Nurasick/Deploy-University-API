FROM golang:1.25-alpine

RUN apk add --no-cache postgresql-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

ENV PORT=8080

EXPOSE 8080

CMD ["./main"]