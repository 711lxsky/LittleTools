# 使用官方的 Go 镜像作为构建环境
FROM golang:1.22.7 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY docker .

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main

# 使用官方的 Alpine Linux 镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8228

# 运行应用程序
CMD ["./main"]
