FROM golang:1.14.2-alpine AS go_build

WORKDIR /go/src/github.com/mf-sakura/docker/echo_server

COPY . /go/src/github.com/mf-sakura/docker/echo_server/

RUN go build -o myapp

# multi-stage build
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM alpine:3.11.6

ENV TZ Asia/Tokyo

# 必要に応じてライブラリをインストールする。不要なものは入れない
# RUN apk --update --no-cache add git build-base openssh

# COPY
WORKDIR /go/src/github.com/mf-sakura/docker/echo_server
ARG REVISION=default
RUN echo $REVISION > .version

COPY --from=go_build /go/src/github.com/mf-sakura/docker/echo_server/myapp .

ENTRYPOINT ["./myapp"]
