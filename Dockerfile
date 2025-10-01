# 使用轻量级的 alpine 镜像作为运行环境
FROM alpine:latest  

# 安装 CA 证书
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# 从构建阶段复制编译好的二进制文件
COPY output/ .

# 暴露应用端口(根据您的应用需要修改)
EXPOSE 8080

# 运行应用
CMD ["./clashmerge"]
