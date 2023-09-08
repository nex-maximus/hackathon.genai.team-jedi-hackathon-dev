########################################################################
 # Copyright (c) Intel Corporation 2023
########################################################################

#build stage
FROM golang:1.20-alpine3.17 AS builder
ARG ALPINE_PKG_BASE="make git gcc libc-dev libsodium-dev zeromq-dev"
ARG ALPINE_PKG_EXTRA=""

RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /app

COPY go.mod ./
RUN go mod download all
COPY . .

ARG MAKE="make build"
RUN $MAKE

#final stage
FROM alpine:3.17
LABEL copyright='Copyright (c) 2023: Intel Corporation' \
      name='app-sample-service' \
      version="${VERSION}"

# dumb-init is required as security-bootstrapper uses it in the entrypoint script
RUN apk add --update --no-cache ca-certificates dumb-init

# COPY --from=builder /app/Attribution.txt /Attribution.txt
COPY --from=builder /app/LICENSE /LICENSE
COPY --from=builder /app/res/ /res/
COPY --from=builder /app/app-sample-service /app-sample-service

EXPOSE 59789

ENTRYPOINT ["/app-sample-service"]
CMD ["-cp=consul.http://edgex-core-consul:8500", "--registry", "--configDir=/res", "-s"]