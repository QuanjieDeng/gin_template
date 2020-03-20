
package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gvp/pkg/app"
	"gvp/pkg/e"
	"gvp/pkg/util"
	log "gvp/pkg/logging"
	"gvp/pkg/setting"
	"gvp/service/game_service"
	"net/http"
)


// @Summary 获取单个游戏信息
// @Produce  json
// @Param   id path int true  "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/game/{id}  [get]
func GetGame(c * gin.Context){
	log.Debug("enter func GetGame")
	appG:=  app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	game_svc :=  game_service.GameService{
		ID:id,
	}

	exist,err :=  game_svc.ExistByID()
	if  err  !=  nil{
		log.Error("GetGame ExistByID:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME_FAIL,nil)
		return
	}
	if  !exist{
		log.Error("GetGame ExistByID:",id,"not exist")
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_GAME,nil)
		return
	}

	svc,err :=   game_svc.Get()
	if err != nil {
		log.Error("GetGame get svc by id:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_GAME_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, svc)
}




// @Summary 获取多个游戏信息
// @Produce  json
// @Param  state query int fase  "账号状态0 锁定 1 正常" Enums(0,1)
// @Param  game query int fase  "游戏名称"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/game  [get]
func GetGames(c * gin.Context){
	log.Debug("enter func GetGames")
	appG:=  app.Gin{c}

	game :=  c.Query("game")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	valid := validation.Validation{}

	valid.MaxSize(game,100,"game").Message("名称最大不能超过100")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}


	game_svc :=  game_service.GameService{
		Game:game,
		State:state,
		PageNum:util.GetPage(c),
		PageSize:setting.AppSetting.PageSize,
	}

	total,err := game_svc.Count()
	if err !=  nil{
		log.Error("GetGames get svc-list by game:",game,"state:",state,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_COUNT_GAME_FAIL,nil)
		return
	}

	svcs,err :=  game_svc.GetAll()
	if err !=  nil{
		log.Error("GetGames get svc-list by game:",game,"state:",state,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_GET_GAME_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,map[string]interface{}{
		"lists":svcs,
		"total":total,
	})
}


type AddGameForm struct {
	Game      	string  `form:"game" valid:"Required;MaxSize(100)"`
	GameType    int 	`form:"gametype" valid:"Required;Range(1,7)"`
	TelNum      string  `form:"telnum" valid:"Required;Phone"`
	CreatedBy 	string  `form:"createdby" valid:"Required;MaxSize(100)"`
	State     	int     `form:"state" valid:"Required;Range(0,1)"`
}

// @Summary 添加游戏
// @Produce  json
// @Param game formData string  true  "游戏名称" maxLength(100)
// @Param gametype formData int true  "游戏类型" Enums(1,7)
// @Param telnum formData string true  "手机号码" maxLength(100)
// @Param createdby formData string true  "创建人" maxLength(100)
// @Param state formData int true  "状态" Enums(0,1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/game  [post]
func AddGame(c * gin.Context){
	var (
		appG = app.Gin{C: c}
		form AddGameForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Error("AddGame  BindAndValid  err:",errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	game_svc := game_service.GameService{
		Game:form.Game,
		GameType:form.GameType,
		TelNum:form.TelNum,
		CreatedBy:form.CreatedBy,
		State:form.State,
	}

	exist,err:=  game_svc.ExistByName()
	if err !=  nil{
		log.Error("AddGame  ExistByName  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME_FAIL,nil)
		return
	}
	if exist{
		log.Error("AddGame   svc   exist:",game_svc.Game)
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME,nil)
		return
	}

	err=   game_svc.Add()
	if  err !=  nil{
		log.Error("AddGame    Add  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_ADD_GAME_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,nil)
}

type EditGameForm struct {
	ID          int     `form:"id" valid:"Required;Min(1)"`
	Game      	string  `form:"game" valid:"MaxSize(100)"`
	GameType    int 	`form:"gametype" valid:"Range(1,7)"`
	TelNum      string  `form:"telnum" valid:"Phone"`
	ModifiedBy  string  `form:"modifiedby" valid:"Required;MaxSize(100)"`
	State     	int     `form:"state" valid:"Range(0,1)"`
}
// @Summary 修改游戏信息
// @Produce  json
// @Param id path int  true  "游戏ID" mininum(1)
// @Param game body string  true  "游戏名称" maxLength(100)
// @Param gametype body int true  "游戏类型" Enums(1,2,3,4,5,6,7)
// @Param telnum body string true  "手机号码" maxLength(100)
// @Param modifiedby body string true  "本次修改人" maxLength(100)
// @Param state body string false  "状态" Enums(0,1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/game/{id} [put]
func UpdateGame(c * gin.Context){
	log.Debug("enter func UpdateGame")

	var (
		appG = app.Gin{C: c}
		form = EditGameForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Error("UpdateGame    BindAndValid  errcode:",errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	game_svc :=game_service.GameService{
		ID:form.ID,
		Game:form.Game,
		GameType:form.GameType,
		TelNum:form.TelNum,
		ModifiedBy:form.ModifiedBy,
		State:form.State,
	}

	exist,err:=  game_svc.ExistByID()
	if err !=  nil{
		log.Error("UpdateGame    ExistByID  errcode:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME_FAIL,nil)
		return
	}
	if !exist{
		log.Error("UpdateGame     svc  not  exist:",game_svc.ID)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_GAME,nil)
		return
	}

	err=   game_svc.Edit()
	if  err !=  nil{
		log.Error("UpdateGame   Edit  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EDIT_GAME_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,nil)
}


// @Summary 删除游戏
// @Produce  json
// @Param id path int  true  "游戏ID" mininum(1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/game/{id} [delete]
func DelGame(c* gin.Context){
	log.Debug("enter func DelGame")
	appG:= app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if  valid.HasErrors(){
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
		return
	}
	game_svc := game_service.GameService{
		ID:id,
	}

	exit,err  := game_svc.ExistByID()
	if  err!=  nil{
		log.Error("DelGame   ExistByID  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_GAME_FAIL,nil)
		return
	}
	if  !exit{
		log.Error("DelGame    svc  not exist:",game_svc.ID)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_VOICESVC,nil)
		return
	}

	err =game_svc.Delete()
	if  err!=  nil{
		log.Error("DelGame    Delete  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_VOICESVC_FAIL,nil)
	}

	appG.Response(http.StatusOK,e.SUCCESS,nil)
}

