#!/bin/sh

# 使用方法: ./wait-for-it.sh host:port command-to-execute

# 检查参数数量
if [ $# -lt 2 ]; then
  echo "Usage: $0 host:port command-to-execute [args...]"
  exit 1
fi

# 解析主机和端口
HOST=$1
PORT=$(echo $HOST | cut -d ':' -f 2)
HOST=$(echo $HOST | cut -d ':' -f 1)

# 检查是否提供了命令
if [ -z "$1" ]; then
  echo "No command to execute provided."
  exit 1
fi

# 等待主机和端口可连接
echo "Waiting for $HOST:$PORT"
while ! nc -z $HOST $PORT; do
  sleep 1
done

shift

# 执行命令
echo "$HOST:$PORT is available. Executing command."
exec "$@"