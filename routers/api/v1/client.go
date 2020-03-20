package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gvp/pkg/app"
	"gvp/pkg/e"
	log "gvp/pkg/logging"
	"gvp/service/client_service"
	"net/http"
)


// @Summary 获取服务商信息
// @Produce  json
// @Param   appid  query string true  "游戏用户appid"
// @Param   appkey query string true  "游戏用户appkey"
// @Param   svcname   query string true  "服务商名称"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/client/svc  [get]
func CliGetVoiceService(c * gin.Context){
	log.Debug("enter func CliGetVoiceService")
	appG:=  app.Gin{c}

	appid  :=  c.Query("appid")
	appkey :=  c.Query("appkey")
	svcname :=  c.Query("svcname")

	valid := validation.Validation{}
	valid.Required(appid,"appid")
	valid.Required(appkey,"appkey")
	valid.Required(svcname,"svcname")
	valid.MaxSize(appid,100,"appid")
	valid.MaxSize(appkey,100,"appkey")
	valid.MaxSize(svcname,100,"svcname")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	cli_svc :=  client_service.Client{
		AppID:appid,
		AppKey:appkey,
		VoiceSvcName:svcname,
	}

	//存在性检查
	exit,err :=  cli_svc.Exist()
	if  err  !=  nil{
		log.Error("CliGetVoiceService  Exist  appid:",appid,"appkey:",appkey,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME_FAIL,nil)
		return
	}
	if  !exit{
		log.Error("CliGetVoiceService  game  not exist appid:",appid,"appkey:",appkey)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_GAME,nil)
		return
	}
	//鉴权
	check,err :=  cli_svc.Check()
	if  err  !=  nil{
		log.Error("CliGetVoiceService  Check  appid:",appid,"appkey:",appkey,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_GAME_FAIL,nil)
		return
	}
	if  !check{
		log.Error("CliGetVoiceService  game   auth false appid:",appid,"appkey:",appkey)
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_GAME_FAIL,nil)
		return
	}
	//获取
	svc,err:= cli_svc.GetVoiceService()
	if err!=  nil{
		log.Error("CliGetVoiceService  GetVoiceService:",appid,"appkey:",appkey,"error:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_GETSVC_GAME_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,svc)


}
