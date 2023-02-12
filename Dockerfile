FROM golang:1.18-alpine as builder

ENV GO115MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /app

COPY . .


RUN go mod download  \
    && go build -o gin-api main.go  \
    && cp /usr/local/go/lib/time/zoneinfo.zip /app/zoneinfo.zip


FROM alpine



## 从builder镜像中把go_app拷贝到当前目录下
WORKDIR /app
COPY --from=builder /app/gin-api .
COPY --from=builder /app/zoneinfo.zip .
COPY app.yml .

ENV ZONEINFO=/app/zoneinfo.zip


# 声明端口
EXPOSE 8080

# 启动容器的命令
CMD ["/app/gin-api"]