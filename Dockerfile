FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/shopus
COPY . .

RUN go build -o ./bin/shopus ./cmd/shopus

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/shopus/bin/shopus /app/shopus

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/shopus

WORKDIR /app
ENTRYPOINT ["/app/shopus"]
