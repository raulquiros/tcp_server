
FROM golang:1.16-buster as build

RUN mkdir -p $GOPATH/src/tcp_server

COPY resources/docker/goLang/start.sh /start.sh
RUN chmod +x /start.sh

ENTRYPOINT ["/start.sh"]
