# yuki-image
## 一个轻量化易部署的图床服务
----------
# API文档
https://apifox.com/apidoc/shared-62b5e13f-632a-4e0a-9e4a-cde4f293e29a
----------
# 可执行文件部署
## 创建一个文件夹，在该文件夹内创建或下载文件config.yaml，内容为
```
server:
    port:  7415                                                             # 服务器端口
    token: XIv3ybWOTIR2Md3sKuMk6AgqjBUH48IRK2d9RMqHGeVymDwc9AWMFOWV7lXc3foJ # 令牌
db:
    max_open_conns: 10                                                      # 最大连接数
    max_idle_conns: 5                                                       # 最大空闲连接数
    reset: true                                                             # 是否重置数据库
image:
    key_length: 8                                                          #key长度
    max_size: 20                                                           # 图片最大大小(MB)
    path: ./localimage                                                     # 图片存储路径
    auto_delete_tmp: true                                                  # 是否自动删除临时文件
    image_list_defalut_size: 10                                            # 图片列表默认大小
    compression_quality: 6                                                 # 图片压缩质量(1-6, 1最低质量, 质量越低速度越快)
    url: http://127.0.0.1:7415                                             # 返回图片的url前缀
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
            - ./data/image:/app/localimage
            - ./data:/app/database
        ports:
            - 7415:7415
```
## 执行
```
docker-compose -p yuki-image up -d
```
