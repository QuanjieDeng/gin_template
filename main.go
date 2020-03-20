package   main


import (
	"fmt"
	"gvp/models"
	log   "gvp/pkg/logging"
	"syscall"
	"github.com/fvbock/endless"
	"gvp/pkg/setting"
	"gvp/routers"
)



// @title GVP API
// @version 0.0.1
// @description  gvp是语音平台的后台管理接口服务器
func main() {

	setting.Setup()
	models.Setup()
	log.Setup()


	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Info("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Info("Server err: %v", err)
	}
}