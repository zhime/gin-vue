FROM golang:1.19.7-alpine3.17 as builder
RUN adduser -u 1000 -D app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /opt/app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o main .


FROM alpine:3.17 as final
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache tzdata
ENV TZ="Asia/Shanghai"
RUN apk update && apk add --no-cache tzdata bash && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && rm -rf /var/lib/apk/*
WORKDIR /opt/app
COPY --from=builder /opt/app/main .
EXPOSE 8090
CMD ["/opt/app/main"]