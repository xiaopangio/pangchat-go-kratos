# pangchat-go-kratos
## 简介
pangchat是一款仿微信的聊天应用，但是也不仅只是一个聊天工具，未来会加入一些更加使用的功能，比如说，听音乐，看电影，便签，备忘录等等。
暂时实现的功能有基本的用户注册，登录，个人信息的维护、好友信息的管理、以及好友之间单聊，未读消息的维护。
接下来要实现的功能就是群聊，以及朋友圈功能，音乐功能。
## 项目依赖
- 项目依赖于docker构建运行环境，需提前安装docker，docker-compose
- 项目用go 1.20开发，请提前安装go 1.20.1
- 前端使用yarn构建，请提前安装yarn
## 项目搭建
### 环境搭建
```
cd build/docker #进入docker目录
docker-compose up -d #docker-compose启动
#进入mysql容器，搭建库环境
pangchat.sql为sql源文件。
```
### 服务配置
source/server 目录下为各服务的源码，需要进入到每个服务代码中，将config.yaml配置好，例如
```
cd source/server/api-gateway/configs
cp config.example.yaml config.yaml 
vim config.yaml 
........
```
### 服务启动
```
cd source/server
make build 会打包所有服务，放在bin目录下。
make run 会打包并运行所有服务。
make stop 会停止所有服务。
make start 会重启所有服务。
```
### 前端构建
前端使用vite构建
```
yarn add 
vite dev 使用开发环境运行
vite build 进行打包
```




