# hadolint global ignore=DL3018
FROM golang:1.23.4-alpine3.20 AS build-env

ARG DB_BACKEND=goleveldb
ARG ROCKSDB_VERSION="9.8.4"

WORKDIR /go/src/github.com/AizelNetwork/evmos

COPY go.mod go.sum ./

RUN set -eux; apk add --no-cache \
    ca-certificates \
    build-base \
    git \
    linux-headers \
    bash \
    binutils-gold

RUN --mount=type=bind,target=. --mount=type=secret,id=GITHUB_TOKEN \
    git config --global url."https://$(cat /run/secrets/GITHUB_TOKEN)@github.com/".insteadOf "https://github.com/"; \
    go mod download

COPY . .

RUN mkdir -p /target/usr/lib /target/usr/local/lib /target/usr/include

RUN if [ "$DB_BACKEND" = "rocksdb" ]; then \
   make build-rocksdb; \
   cp -r /usr/lib/* /target/usr/lib/ && \
   cp -r /usr/local/lib/* /target/usr/local/lib/ && \
   cp -r /usr/include/* /target/usr/include/; \
else \
    # Build default binary with corresponding db backend
    COSMOS_BUILD_OPTIONS=$DB_BACKEND make build; \
fi

RUN go install github.com/MinseokOh/toml-cli@latest

FROM alpine:3.21

WORKDIR /root

COPY --from=build-env /go/src/github.com/AizelNetwork/evmos/build/aizeld /usr/bin/aizeld
COPY --from=build-env /go/bin/toml-cli /usr/bin/toml-cli

# required for rocksdb build
COPY --from=build-env /target/usr/lib /usr/lib
COPY --from=build-env /target/usr/local/lib /usr/local/lib
COPY --from=build-env /target/usr/include /usr/include

RUN apk add --no-cache \
    ca-certificates \
    jq \
    curl \
    bash \
    vim \
    lz4 \
    rclone \
    && addgroup -g 1000 aizel \
    && adduser -S -h /home/aizel -D aizel -u 1000 -G aizel

USER 1000
WORKDIR /home/aizel

EXPOSE 56656 56657 11317 19090 18545 18546
HEALTHCHECK CMD curl --fail http://localhost:56657 || exit 1

CMD ["aizeld"]
