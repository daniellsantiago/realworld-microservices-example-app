FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

FROM gcr.io/distroless/base

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]