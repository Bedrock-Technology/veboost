ARG BASE_REGISTRY=docker.io/library

FROM ${BASE_REGISTRY}/golang:1.27.4-alpine AS builder

WORKDIR /src
ENV CGO_ENABLED=0

COPY merkle/go.mod merkle/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY merkle/ ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go test ./... && \
    go build -mod=readonly -trimpath -ldflags="-s -w" -o /out/VeMerkle .

FROM ${BASE_REGISTRY}/alpine:3.24

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S -g 10001 vemerkle && \
    adduser -S -D -H -u 10001 -G vemerkle vemerkle && \
    mkdir -p /app/configs /app/docs /app/logs /app/data && \
    chown -R vemerkle:vemerkle /app

WORKDIR /app
COPY --from=builder /out/VeMerkle /usr/local/bin/VeMerkle
COPY --from=builder /src/docs/swagger.json ./docs/swagger.json

USER vemerkle
EXPOSE 8080

ENTRYPOINT ["VeMerkle"]
CMD ["run", "--port", "8080"]
