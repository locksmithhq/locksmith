FROM alpine:3.20 AS timezone_build_api
RUN apk --no-cache add tzdata ca-certificates  

FROM golang:1.25.5-alpine3.21 AS builder_api
RUN apk --no-cache add tzdata ca-certificates

WORKDIR /go/api

COPY ./api .

RUN mkdir deploy
RUN go clean --modcache
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o locksmith cmd/main.go 
RUN mv locksmith ./deploy/locksmith
ADD ../config/database ./deploy/database

FROM oven/bun:1 AS builder_web

WORKDIR /web

COPY ./web . 

COPY ./web/package.json ./web/bun.lock ./

RUN bun install

RUN bun run build

FROM scratch AS production

COPY --from=timezone_build_api /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=timezone_build_api /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder_api /go/api/deploy /api/
COPY --from=builder_web /web/dist /web/dist
COPY ../config/database/seeder /etc/locksmith/config

WORKDIR /api

ENTRYPOINT  ["./locksmith", "serve"]