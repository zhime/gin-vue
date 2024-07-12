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
WORKDIR /opt/app
COPY --from=builder /opt/app/main .
EXPOSE 8080
USER 1000:1000
CMD ["/opt/app/main"]