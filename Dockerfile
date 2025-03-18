FROM golang:1.24.1-alpine AS builder
COPY . /go/src/rubik-web
WORKDIR /go/src/rubik-web
RUN go mod download
RUN go build -o ./bin/web_server cmd/web_server/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/rubik-web/bin/web_server .
COPY .env .
CMD ["./web_server"]
