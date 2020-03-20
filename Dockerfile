FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/gvp
COPY . $GOPATH/src/gvp
RUN go build  -mod=vendor  .

EXPOSE 8000
ENTRYPOINT ["./gvp"]