FROM alpine:latest

LABEL MAINTAINER="soul.sxd@gmail.com"

RUN mkdir /app/

ADD app /app/

WORKDIR /app

EXPOSE 8088

CMD ["./app"]
