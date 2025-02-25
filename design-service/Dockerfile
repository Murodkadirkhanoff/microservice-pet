FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app

FROM alpine

WORKDIR /

COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
