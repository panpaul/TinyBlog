FROM golang:alpine as server
WORKDIR /app
COPY server /app
RUN apk add --no-cache make
RUN make all

FROM node:16-alpine as frontend
WORKDIR /app
COPY frontend /app
RUN yarn
RUN yarn run build

FROM alpine:latest
LABEL MAINTAINER="panyuxuan@hotmail.com"
WORKDIR /app
COPY --from=server /app/server /app/
COPY --from=server /app/resource /app/resource/
COPY --from=server /app/config.docker.yaml /app/
COPY --from=frontend /app/build /app/resource/frontend/
COPY docker-entrypoint.sh /app/
EXPOSE 8080
ENTRYPOINT ["/app/docker-entrypoint.sh"]
CMD ["run"]
