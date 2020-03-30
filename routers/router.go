package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	log  "gvp/pkg/logging"
	_ "gvp/docs"
	"gvp/pkg/setting"
	"gvp/routers/api/v1"
	"gvp/routers/api"
	"gvp/middleware/jwt"
)

type resBody struct {

}


func InitRouter() *gin.Engine {
	log.Info("初始化路由!")
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	if  setting.ServerSetting.RunMode  == "debug"{
		gin.SetMode(gin.DebugMode)
	}else if setting.ServerSetting.RunMode  == "release"{
		gin.SetMode(gin.ReleaseMode)
	}

	//鉴权
	r.GET("/auth", api.GetAuth)

	//文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 :=  r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{


		//获取单个服务商信息
		apiv1.GET("/svc/:id", v1.GetVoiceService)
		//获取服务商列表
		apiv1.GET("/svc", v1.GetVoiceServices)
		//添加服务商
		apiv1.POST("/svc", v1.AddVoiceService)
		//更新指定的服务商
		apiv1.PUT("/svc/:id", v1.UpdateVoiceService)
		//删除指定的服务商
		apiv1.DELETE("/svc/:id", v1.DelVoiceService)




		//获取游戏列表
		apiv1.GET("/game", v1.GetGames)
		//获取指定游戏
		apiv1.GET("/game/:id", v1.GetGame)
		//新增游戏
		apiv1.POST("/game", v1.AddGame)
		//更新游戏信息
		apiv1.PUT("/game/:id", v1.UpdateGame)
		//删除指定游戏
		apiv1.DELETE("/game/:id", v1.DelGame)


		//获取绑定关系列表
		apiv1.GET("/bind", v1.GetBinds)
		//获取单个绑定关系
		apiv1.GET("/bind/:id", v1.GetBind)
		//添加绑定关系
		apiv1.POST("/bind", v1.AddBind)
		//删除绑定关系
		apiv1.DELETE("/bind/:id",v1.DelBind)
	}

	clientv1 :=  r.Group("/api/v1/client")
	clientv1.Use(jwt.JWT())
	{
		//获取服务商信息
		clientv1.GET("/svc",v1.CliGetVoiceService)
	}


	log.Info("路由初始化完成！")

	return r
}