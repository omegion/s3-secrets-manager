ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.14

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.source="https://github.com/omegion/s3-secrets-manager-template"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make curl unzip && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH VERSION=$VERSION

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/s3sm /bin/s3sm

ENTRYPOINT ["s3sm"]
