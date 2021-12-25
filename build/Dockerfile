FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct

WORKDIR /go/src/golang-gin-demo
COPY ./ ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
        apk update && \
        apk add --no-cache git

WORKDIR /go/src/golang-gin-demo/cmd
RUN go env && go build -o golang-gin-demo .

FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /go/src/golang-gin-demo
COPY --from=0 /go/src/golang-gin-demo/cmd ./

EXPOSE 8080
RUN ls

ENTRYPOINT ./golang-gin-demo
