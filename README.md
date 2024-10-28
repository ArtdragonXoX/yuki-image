# yuki-image
## 一个简易的使用本地存储的图床服务
----------
# API文档
https://apifox.com/apidoc/shared-666c1f43-0462-4b14-a115-7627ff0403bc
----------
# 可执行文件部署
##注意：需要mysql
## 创建一个文件夹，在该文件夹内创建或下载文件config.yaml，内容为
```
server:
    port: "7415"                                                            # 服务器端口
    host: http://127.0.0.1                                                  # 服务器地址
    token: XIv3ybWOTIR2Md3sKuMk6AgqjBUH48IRK2d9RMqHGeVymDwc9AWMFOWV7lXc3foJ # 令牌
db:
    host: yuki-image-db                                                     # mysql地址
    port: "3306"                                                            # mysql端口
    name: yuki_image_db                                                     # mysql名称
    user: yuki-image                                                        # mysql用户
    password: yuki-image                                                    # mysql密码
    max_open_conns: 10                                                      # 最大连接数
    max_idle_conns: 5                                                       # 最大空闲连接数
    reset: true                                                             # 是否重置数据库
image:
    max_size: 20                                                           # 图片最大大小(MB)
    path: ./localimage                                                     # 图片存储路径
    image_list_defalut_size: 10                                            # 图片列表默认大小
    compression_quality: 6                                                 # 图片压缩质量(1-6, 1最低质量, 质量越低速度越快)
```
## 在release下载对应可执行文件在同一文件夹，执行
```
./可执行文件名 server
```
------------
# docker compose部署
## 在release下载docker部署压缩包并解压，包含config.yaml和docker-compose.yml
docker-compose.yml文件内容：
```
version: '3'
services:
    yuki-image:
        container_name: yuki-image
        image: crpi-jg7reu90mmam5qch.cn-guangzhou.personal.cr.aliyuncs.com/artdragon/yuki-image
        restart: no
        volumes:
            - ./config.yaml:/app/config.yaml
            - ./data/tmp:/app/tmp
            - ./data/image:/app/localimage
        ports:
            - 7415:7415
    mysql:
        image: mysql:8.0
        container_name: yuki-image-db
        restart: no
        hostname: yuki-image-db
        environment:
          - MYSQL_DATABASE=yuki_image_db
          - MYSQL_USER=yuki-image
          - MYSQL_PASSWORD=yuki-image
          - MYSQL_ROOT_PASSWORD=yuki-image
        volumes:
          - ./data/db:/var/lib/mysql
```
## 执行
```
docker-compose -p yuki-image up -d
```