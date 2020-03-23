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
    ├── gvp
    │   ├── docs     //swag 生成的接口文件说明
    │   │   ├── docs.go
    │   │   ├── swagger.json
    │   │   └── swagger.yaml
    │   ├── middleware   //中间件
    │   │   └── jwt
    │   │       └── jwt.go   
    │   ├── models      //数据模型
    │   │   ├── auth.go
    │   │   ├── models.go
    │   ├── pkg        //公共包
    │   │   ├── e 错误
    │   │   │   ├── code.go    //错误码
    │   │   │   └── msg.go     //错误码说明
    │   │   ├── logging 日志       
    │   │   │   ├── file.go
    │   │   │   └── log.go
    │   │   ├── setting 配置
    │   │   │   └── setting.go
    │   │   └── util 功能函数
    │   │       ├── jwt.go
    │   │       └── pagination.go
    │   └── routers 路由
    │       ├── api
    │       │   ├── auth.go
    │       │   └── v1  v1版本 
    │       └── router.go
       main.go  入口文件
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



