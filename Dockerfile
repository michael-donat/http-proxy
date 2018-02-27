FROM golang:1.10-alpine AS build

RUN apk update && apk add make git gcc musl-dev

ADD . /build

WORKDIR /build

RUN go build -o http-proxy

FROM alpine:3.7

RUN apk add --no-cache ca-certificates

COPY --from=build /build/http-proxy /http-proxy

ENTRYPOINT ["/http-proxy"]