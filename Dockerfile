FROM golang:1.12.9-alpine3.10 as builder

WORKDIR /gitbook-api
ADD . .


ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0

RUN ls -al && \
    go env && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache make && \
    make build

FROM alpine:latest

LABEL MAINTAINER="soul.sxd@gmail.com"

RUN mkdir -p /app/config

COPY --from=builder /gitbook-api/gitbook-api /app/

WORKDIR /app

ENV RUNTIME_CONFIG=/app/config

EXPOSE 8088

CMD ["./gitbook-api"]
