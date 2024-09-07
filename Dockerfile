FROM golang:1.21.1-alpine AS build

RUN apk update && apk add --no-cache make

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make docker

# Release stage
FROM alpine:3.18

RUN apk update && apk add --no-cache chromium

WORKDIR /bin
COPY --from=build /usr/local/bin/igdb-app /bin/igdb-app
RUN touch .env
RUN chmod +x /bin/igdb-app

CMD ["igdb-app", "-headless"]