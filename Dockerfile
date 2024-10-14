# 使用官方 Go 镜像作为构建环境
FROM golang:1.22 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o clashmerge .


# 使用轻量级的 alpine 镜像作为运行环境
FROM alpine:latest  

# 安装 CA 证书
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/clashmerge .

# 暴露应用端口(根据您的应用需要修改)
EXPOSE 8080

# 运行应用
CMD ["./clashmerge"]
