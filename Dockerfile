FROM golang:1.18-alpine AS builder

ARG APP_ENV=local
ENV APP_ENV ${APP_ENV}
CMD [ "echo", "building environment ${APP_ENV}" ]

WORKDIR /src
COPY go.mod .
COPY go.sum .
# COPY ${APP_ENV}.env .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main /src/cmd/server/grpc/main.go

FROM alpine:latest

# Binary
COPY --from=builder /src /app

# Set app environment as local or override during build ie.:
# $ docker build --build-arg APP_ENV=prod .
# ARG APP_ENV=local
# ENV APP_ENV ${APP_ENV}
# COPY --from=builder /src/${APP_ENV}.env .
# Go certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

ENTRYPOINT [ "/app/main" ]
