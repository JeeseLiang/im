# 第一阶段：构建
FROM golang:1.23.9-alpine as builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制整个项目
COPY . .

# 构建应用
RUN go build -o msg-rpc app/msg/rpc/msg.go

# 第二阶段：运行
FROM alpine:latest

# 安装必要的依赖
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 复制配置文件
COPY app/msg/rpc/etc /app/etc

# 从builder阶段复制构建好的应用
COPY --from=builder /app/msg-rpc /app/

# 暴露端口
EXPOSE 9002

# 启动应用
CMD ["./msg-rpc", "-f", "etc/msg-rpc.yaml"]
