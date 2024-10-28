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
  port: 7415                  #服务端口号
  path: ./localimage          #图片本地存储路径
  host: http://127.0.0.1      #图片url前缀中的host
  image_list_defalut_size: 10 #查询图片列表默认图片数量
  token: XIv3ybWOTIR2Md3sKuMk6AgqjBUH48IRK2d9RMqHGeVymDwc9AWMFOWV7lXc3foJ  #认证token，请替换为示例值
db:
  host: yuki-image-db    #mysql主机host
  port: 3306             #mysql端口号
  name: yuki_image_db    #mysql库名称
  user: yuki-image       #mysql用户名
  password: yuki-image   #mysql密码
  max_open_conns: 10     #mysql最大连接数
  max_idle_conns: 5      #mysql保持连接数
  reset: true            #重置数据库标志
```
## 在release下载对应可执行文件在同一文件夹，执行
｀｀｀
./可执行文件名 server
｀｀｀
------------
# docker compose部署
## 在release下载docker部署压缩包并解压，包含config.yaml和docker-compose.yml
docker-compose.yml文件内容：
｀｀｀
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
｀｀｀
## 执行
｀｀｀
docker-compose -p yuki-image up -d
｀｀｀