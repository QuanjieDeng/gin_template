# GVP

语音平台-后台
- 开发模式为WEB常规MVC模式
- 基于Gin框架开发的API服务器



## 接口文档 

- API文档为在线文档，使用swag生成  
- 确保程序正常运行，然后访问路径为http://IP:PORT/swagger/index.html
- 更新：当需要更新文档时 在项目根目录执行： swag  init 即可


## 项目结构说明

```
├── conf   配置文件存放处
│   └── app.ini
├── cron  定时任务
│   ├── cron.go
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│   └── vendor
│       ├── github.com
│       │   └── robfig
│       │       └── cron
│       │           ├── constantdelay.go
│       │           ├── cron.go
│       │           ├── doc.go
│       │           ├── LICENSE
│       │           ├── parser.go
│       │           ├── README.md
│       │           └── spec.go
│       └── modules.txt
├── Dockerfile   生成镜像的DockerFile
├── docs   swager生成的DOC文件
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go    程序主入口
├── Makefile   
├── middleware  中间件
│   └── jwt
│       └── jwt.go
├── models     数据模型
│   ├── auth.go
│   ├── game.go
│   ├── gameServiceRelation.go
│   ├── models.go
│   └── voiceService.go
├── pkg      工具包
│   ├── app
│   │   ├── form.go
│   │   ├── request.go
│   │   └── response.go
│   ├── e
│   │   ├── code.go
│   │   ├── gametype.go
│   │   └── msg.go
│   ├── file
│   │   └── file.go
│   ├── logging
│   │   ├── file.go
│   │   └── log.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── README.md
├── routers   路由规划
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── bindmabager.go
│   │       ├── client.go
│   │       ├── game.go
│   │       └── voiceservice.go
│   └── router.go
├── runtime
│   └── logs
├── service    服务封装
│   ├── auth_service
│   │   └── auth.go
│   ├── bind_service
│   │   └── bind.go
│   ├── client_service
│   │   └── client.go
│   ├── game_service
│   │   └── game.go
│   └── voicesvc_service
│       └── voiceservice.go
├── sql
│   └── init.sql
├── test    单元测试文件夹
│   ├── auth_test.go
│   ├── bind_test.go
│   ├── conf
│   │   └── app.ini
│   ├── game_test.go
│   ├── runtime
│   │   └── logs
│   └── voiceservice_test.go
```


## How to Start

### 依赖组件

- mysql 
- docker  

### 配置

- 配置文件存放在 conf/conf.ini  根据实际情况进行配置

### 编译

- 项目基于go  mod模式开发，确保 go mod 打开
- 切换到项目跟目录

``
make 
``

### 数据库初始化

- sql/init.sql 为初始化数据库脚本，通过mysql执行即可完成数据库的初始化

### 运行

- 完成编译和数据初始化后，直接执行生成 gvp可执行文件即可

## 接口访问鉴权

- 统一使用jwt鉴权，访问者需要先使用username+passwd 申请token token值在一段时间后会过期  可以通过AuthExpire参数配置 默认是5小时
- 在随后的请求中将得到的token作为URL参数的形式附加在每次请求中
### 请求token的接口

```
http://192.168.94.145:8000/auth?username=admin&password=admin
```
### 接口请求格式
```
http://192.168.94.145:8000/api/v1/bind?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbiIsImV4cCI6MTU4NDcwMjA4MCwiaXNzIjoiZ2luLWJsb2cifQ.mzb0Rh104wntMTySy2SjVnE2WaKm9C7WV6NoFRShfug
```

## Docker

- 由于项目是在go  mod 模式下开发，在生成镜像文件时采取的是直接指定的vendor文件夹，所以需要事先生成
go mod  vendor  
- docker    build  -t  avp  生成镜像文件
- go run  -p 8000:8000  avp   运行时做端口映射



## K8S
- 如果需要部署在K8S环境中，相关的部署文件存放在 K8SPlateFrom 文件夹下面，参考其中的说明文档即可

## Unit-Test  

### 说明

- 测试框架 httptest+Convery + Stub
- 测试代码在test文件夹下
- conf文件夹是测试项目的配置文件，他是直接从项目的conf文件夹拷贝来的，由于测试需要一些特殊的配置所以单独使用了一个配置文件
- 测试使用单独的DB，建议创建单独的数据库gvp_test，然后在conf/app.ini的配置文件中修改[database]Name配置项
- 在创建好DB后，运行sql/init.sql文件创建数据库表
- 运行测试，切换到test目录执行以下命令
```
    goconvey   -host="0.0.0.0" -port=8090    
```
- 在浏览器中访问  http://youri-p:8090  即可看到Convery提供的可视化测试界面


## 其他

## 使用Nginx部署

- nginx  -t 查看nginx的配置文件位置 
- 修改nginx.confg  添加如下配置

```
    upstream api.blog.com {
        server 127.0.0.1:8001;
        server 127.0.0.1:8002;
    }

    server {
        listen       8081;
        server_name  api.blog.com;

        location / {
            proxy_pass http://api.blog.com/;
        }
    }
```

- nginx  -s reload 重新加载配置即可


### mysql的安装 

https://www.cnblogs.com/tianphone/p/10767886.html

### mysql-docker

- 项目需要mysql组件，可以直接在机器中安装mysql 或者是采用mysql容器的方式提供服务
可以参考 https://www.jianshu.com/p/d6febf6f95e0



