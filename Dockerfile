FROM golang:1.21.1-alpine as build

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/updater ./cmd/updater/main.go
RUN go build -v -o /usr/local/bin/watcher ./cmd/watcher/main.go

# Release stage
FROM alpine:3.18

WORKDIR /bin
COPY --from=build /usr/local/bin/updater /bin/updater
COPY --from=build /usr/local/bin/watcher /bin/watcher
COPY .env .