FROM golang:alpine as builder

WORKDIR /app
COPY . /app

RUN apk add --no-cache make

RUN make all

FROM alpine:latest

LABEL MAINTAINER="panyuxuan@hotmail.com"

WORKDIR /app

COPY --from=builder /app/server /app/
COPY --from=builder /app/resource /app/resource/
COPY --from=builder /app/config.docker.yaml /app/

EXPOSE 8080
ENTRYPOINT ./server -c config.docker.yaml run
