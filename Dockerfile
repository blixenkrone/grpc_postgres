FROM golang:1.18-alpine as builder

WORKDIR /src
COPY go.mod .
COPY go.sum .
COPY /conf/ .
COPY .env .

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main /src/cmd/server/main.go

FROM alpine:latest

# Binary
COPY --from=builder /src /app
COPY --from=builder /src/.env /app
# Go certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

ENTRYPOINT [ "/app/main" ]
