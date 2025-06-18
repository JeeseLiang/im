# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的构建依赖
RUN apk add --no-cache git

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译所有服务
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/user/user-rpc ./app/user/rpc/user.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/user/user-api ./app/user/api/user.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/group/group-rpc ./app/group/rpc/group.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/group/group-api ./app/group/api/group.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/msg/msg-rpc ./app/msg/rpc/msg.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/bin/msg/msg-api ./app/msg/api/msg.go

# 运行阶段
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache bash tzdata ca-certificates

# 设置时区
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/cmd/bin/user/user-rpc app/cmd/bin/user/
COPY --from=builder /app/cmd/bin/user/user-api app/cmd/bin/user/
COPY --from=builder /app/cmd/bin/group/group-rpc app/cmd/bin/group/
COPY --from=builder /app/cmd/bin/group/group-api app/cmd/bin/group/
COPY --from=builder /app/cmd/bin/msg/msg-rpc app/cmd/bin/msg/
COPY --from=builder /app/cmd/bin/msg/msg-api app/cmd/bin/msg/

# 复制配置文件
COPY --from=builder /app/app/user/rpc/etc/user-rpc.yaml app/cmd/bin/user/etc/user-rpc.yaml
COPY --from=builder /app/app/user/api/etc/user-api.yaml app/cmd/bin/user/etc/user-api.yaml
COPY --from=builder /app/app/group/rpc/etc/group-rpc.yaml app/cmd/bin/group/etc/group-rpc.yaml
COPY --from=builder /app/app/group/api/etc/group-api.yaml app/cmd/bin/group/etc/group-api.yaml
COPY --from=builder /app/app/msg/rpc/etc/msg-rpc.yaml app/cmd/bin/msg/etc/msg-rpc.yaml
COPY --from=builder /app/app/msg/api/etc/msg-api.yaml app/cmd/bin/msg/etc/msg-api.yaml

# 复制并设置启动脚本cl
COPY start.docker.sh app/start.sh
COPY .env app/.env
RUN chmod +x /app/app/start.sh

# 暴露所有服务端口
# RPC服务端口
EXPOSE 10002 20002 30002
# API服务端口
EXPOSE 10001 20001 30001

# 设置入口点
ENTRYPOINT ["/app/app/start.sh"]