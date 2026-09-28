FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/bin \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/bin/worker ./cmd/worker


FROM alpine:3.22 AS runtime

WORKDIR /app

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S app -G app

COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/bin/worker /app/worker

USER app

CMD ["/app/api"]