# Stage 1: CA certs + timezone data + non-root user definition
FROM alpine:3.21 AS base
RUN apk --no-cache add tzdata ca-certificates && \
    addgroup -S locksmith && adduser -S locksmith -G locksmith

# Stage 2: Go binary
FROM golang:1.25.5-alpine3.21 AS builder_api
WORKDIR /build

# Dependency layer cached separately — only re-runs when go.mod/go.sum change
COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ .
RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -trimpath \
    -o /locksmith \
    cmd/main.go

# Stage 3: Frontend
FROM oven/bun:1-alpine AS builder_web
WORKDIR /build

COPY web/package.json web/bun.lock ./
RUN bun install --frozen-lockfile

COPY web/ .
RUN bun run build

# Stage 4: Minimal production image
FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=builder_api /locksmith /locksmith
COPY config/database/migrations /api/database/migrations
COPY --from=builder_web /build/dist /web/dist
COPY config/database/seeder /etc/locksmith/config

USER locksmith:locksmith

ENTRYPOINT ["/locksmith", "serve"]
