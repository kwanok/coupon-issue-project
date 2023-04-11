FROM golang:1.19-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . .

RUN go mod tidy

RUN go build -o main main.go

FROM alpine:latest

WORKDIR /app

RUN apk --update add ca-certificates

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["/app/main"]