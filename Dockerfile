FROM golang:1.21.2-alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go mod download
RUN go build -o auth cmd/grpc-server/user_v1/main.go
RUN apk update --no-cache && apk add --no-cache ca-certificates

FROM alpine

WORKDIR /app

COPY --from=builder /build/auth .

CMD ["./auth"]