# GVP

语音平台-后台


##编译

- 项目基于go  mod模式开发，确保 go mod 打开

##接口文档 
- API文档为在线文档，使用swag生成  访问路径为http://IP:PORT/swagger/index.html
- 更新：当需要更新文档时 在项目根目录执行： swag  init 即可


##项目结构说明

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


##运行
###依赖组件
- mysql 
- docker  

###配置
- 配置文件存放在 conf/conf.ini  根据实际情况进行配置

###直接运行
- 可执行文件为编译的到 gvp 直接运行即可 

###Docker

- 由于项目是在go  mod 模式下开发，在生成镜像文件时采取的办吧是直接指定的vendor文件夹，所以需要事先生成
go mod  vendor  

- docker    build  -t  avp  生成镜像文件
- go run  -p 8000:8000  avp   运行时做端口映射

