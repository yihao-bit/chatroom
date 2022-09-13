# 聊天室

#### 介绍
多人聊天室，实现了群发，上线提醒功能

可以注册和登录用户

使用 docker 进行部署

# server
把某个用户发送的消息群送到其他用户

服务器得创建 名字叫 users 数据库，名字叫 user 的表，usr表结构为（usrId int, usrPwd string, userName string）

# client

客户端需要把连接服务端的 IP 地址和端口改为自己所要连接的服务器的ip地址和端口。


使用的是 golang，编译需要 golang环境，或者可以用docker部署。





