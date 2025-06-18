#!/bin/bash

# 设置颜色输出
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 定义服务启动函数
start_service() {
    local service_name=$1
    local service_type=$2

    echo -e "${BLUE}Starting ${service_name} ${service_type}...${NC}"
    # 启动服务
    cd /app/app/cmd/bin/${service_name,,} && ./${service_name,,}-${service_type,,} &
    # /app/app/cmd/bin/${service_name,,}/${service_name,,}-${service_type,,} &
    echo -e "${GREEN}${service_name} ${service_type} started${NC}"
    sleep 2
}

# 启动顺序：先启动RPC服务，再启动API服务

# 启动用户服务
start_service "user" "rpc"
start_service "user" "api"

# 启动群组服务
start_service "group" "rpc"
start_service "group" "api"

# 启动消息服务
start_service "msg" "rpc"
start_service "msg" "api"

echo -e "${GREEN}All services started successfully!${NC}"

# 等待所有后台进程
wait 