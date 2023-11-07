FROM golang:1.21-alpine AS builder
LABEL authors="arseniy"

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ./ ./
COPY ./internal/server/storage/db/migrations ./internal/storage/db/migrations/
RUN go build -o ./bin/gophkeeper-server cmd/server/main.go

CMD ["./bin/gophkeeper-server"]