FROM golang:1.12 as builder

WORKDIR /gitbook-api
ADD . .

ENV GOPROXY="https://goproxy.io"

RUN ls -al && \
    go env && \
    make build

FROM alpine:latest

LABEL MAINTAINER="soul.sxd@gmail.com"

RUN mkdir -p /app/config

COPY --from=builder /gitbook-api/gitbook-api /app/

WORKDIR /app

ENV RUNTIME_CONFIG=/app/config

EXPOSE 8088

CMD ["./gitbook-api"]
