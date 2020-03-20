package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	log  "gvp/pkg/logging"
	"gvp/pkg/app"
	"gvp/pkg/e"
	"gvp/pkg/setting"
	"gvp/pkg/util"
	"gvp/service/bind_service"
	"net/http"
)

// @Summary 获取多个绑定关系
// @Produce  json
// @Param  gameid  query int  fase  "游戏名称"
// @Param  serviceid  query int fase  "游戏名称"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/bind  [get]
func GetBinds(c * gin.Context){
	log.Debug("enter func GetBinds")
	appG:=  app.Gin{c}

	var gameid int = -1
	if arg := c.Query("gameid"); arg != "" {
		gameid = com.StrTo(arg).MustInt()
	}
	var serviceid int = -1
	if arg := c.Query("serviceid"); arg != "" {
		serviceid = com.StrTo(arg).MustInt()
	}


	bind_svc :=  bind_service.BindService{
		GameId:gameid,
		ServiceId:serviceid,
		PageNum:util.GetPage(c),
		PageSize:setting.AppSetting.PageSize,
	}

	total,err := bind_svc.Count()
	if err !=  nil{
		log.Error("GetBinds get svc-list by gameid:",gameid,"serviceid:",serviceid,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_COUNT_BIND_FAIL,nil)
		return
	}

	svcs,err :=  bind_svc.GetAll()
	if err !=  nil{
		log.Error("GetBinds get svc-list by gameid:",gameid,"serviceid:",serviceid,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_GET_BIND_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,map[string]interface{}{
		"lists":svcs,
		"total":total,
	})
}




// @Summary 获取单个绑定信息
// @Produce  json
// @Param   id path int true  "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/bind/{id}  [get]
func GetBind(c * gin.Context){
	log.Debug("enter func GetBind")
	appG:=  app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	bind_svc :=  bind_service.BindService{
		ID:id,
	}

	exist,err :=    bind_svc.ExistByID()
	if  err  !=  nil{
		log.Error("GetBind ExistByID:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_BIND_FAIL,nil)
		return
	}
	if  !exist{
		log.Error("GetBind ExistByID:",id,"not exist")
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_BIND,nil)
		return
	}

	svc,err :=  bind_svc.Get()
	if err != nil {
		log.Error("GetBind get svc by id:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_BIND_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, svc)
}


type AddBindForm struct {
	Game      	 string  `form:"game" valid:"Required;MaxSize(100)"`
	ServiceId    int 	 `form:"serviceid" valid:"Required;Min(1)"`
	GameId       int     `form:"gameid" valid:"Required;Min(1)"`
}

// @Summary 添加绑定关系
// @Produce  json
// @Param game formData string  true  "游戏名称" maxLength(100)
// @Param serviceid formData int true  "服务ID"
// @Param gameid formData int true  "游戏ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/bind  [post]
func AddBind(c * gin.Context){
	var (
		appG = app.Gin{C: c}
		form AddBindForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Error("AddBind  BindAndValid  err:",errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	bind_svc := bind_service.BindService{
		Game:form.Game,
		ServiceId:form.ServiceId,
		GameId:form.GameId,
	}

	exist,err:=  bind_svc.ExistBySerIdAndGameId()
	if err !=  nil{
		log.Error("AddBind  ExistByName  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_BIND_FAIL,nil)
		return
	}
	if exist{
		log.Error("AddBind   svc   exist game:",bind_svc.Game,"Service_id:",bind_svc.ServiceId,"gameid:",bind_svc.GameId)
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_BIND,nil)
		return
	}

	err=  bind_svc.Add()
	if  err !=  nil{
		log.Error("AddBind    Add  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_ADD_BIND_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,nil)
}

// @Summary 删除绑定关系
// @Produce  json
// @Param id path int  true  "绑定关系ID" mininum(1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/bind/{id} [delete]
func DelBind(c* gin.Context){
	log.Debug("enter func DelBind")
	appG:= app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if  valid.HasErrors(){
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
		return
	}
	bind_svc := bind_service.BindService{
		ID:id,
	}

	exit,err  := bind_svc.ExistByID()
	if  err!=  nil{
		log.Error("DelBind   ExistByID  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_BIND_FAIL,nil)
		return
	}
	if  !exit{
		log.Error("DelBind    svc  not exist:",bind_svc.ID)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_BIND,nil)
		return
	}

	err =bind_svc.Delete()
	if  err!=  nil{
		log.Error("DelGame    Delete  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_BIND_FAIL,nil)
	}

	appG.Response(http.StatusOK,e.SUCCESS,nil)
}
