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
    local service_path=$3
    local main_file=$4

    echo -e "${BLUE}Starting ${service_name} ${service_type}...${NC}"
    cd "${service_path}" || exit
    # 启动服务
    cmd.exe /c "go run ${main_file}" &
    cd - > /dev/null || exit
    echo -e "${GREEN}${service_name} ${service_type} started${NC}"
    sleep 2
}

# 确保脚本在项目根目录运行
if [ ! -d "app" ]; then
    echo "Please run this script from the project root directory"
    exit 1
fi

# 启动顺序：先启动RPC服务，再启动API服务

# 启动用户服务
start_service "User" "rpc" "app/user/rpc" "user.go"
start_service "User" "api" "app/user/api" "user.go"

# 启动群组服务
start_service "Group" "rpc" "app/group/rpc" "group.go"
start_service "Group" "api" "app/group/api" "group.go"

# 启动消息服务
start_service "Message" "rpc" "app/msg/rpc" "msg.go"
start_service "Message" "api" "app/msg/api" "msg.go"

echo -e "${GREEN}All services started successfully!${NC}"

# 等待所有后台进程
wait 