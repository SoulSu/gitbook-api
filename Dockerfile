FROM alpine:latest

LABEL MAINTAINER="soul.sxd@gmail.com"

RUN mkdir -p /app/config

ADD app /app/

WORKDIR /app

ENV RUNTIME_CONFIG=/app/config

EXPOSE 8088

CMD ["./app"]
