FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR   /opt/src/gvp
COPY .  /opt/src/gvp
RUN go build  -mod=vendor  .

ENTRYPOINT ["./gvp"]
