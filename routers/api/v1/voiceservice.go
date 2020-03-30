package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	log  "gvp/pkg/logging"
	"gvp/pkg/e"
	"gvp/pkg/setting"
	"gvp/pkg/util"
	"net/http"
	"gvp/pkg/app"
	"gvp/service/voicesvc_service"

)

// @Summary 获取单个服务商信息
// @Produce  json
// @Param   id path int true  "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/svc/{id}  [get]
func GetVoiceService(c * gin.Context){
	log.Debug("enter func GetVoiceService")
	appG:=  app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}


	voice_svc :=  voicesvc_service.VoiceService{
		ID:id,
	}
	exist,err :=  voice_svc.ExistByID()
	if  err  !=  nil{
		log.Error("GetVoiceService ExistByID:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_VOICESVCFAIL,nil)
		return
	}
	if  !exist{
		log.Error("GetVoiceService ExistByID:",id,"not exist")
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_VOICESVC,nil)
		return
	}

	svc,err :=   voice_svc.Get()
	if err != nil {
		log.Error("GetVoiceService get svc by id:",id,"err:",err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VOICESVC_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, svc)
}




// @Summary 获取多个服务商信息
// @Produce  json
// @Param state query int fase  "服务状态0 锁定 1 正常" Enums(0,1)
// @Param  name query int fase  "服务商名称"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/svc  [get]
func GetVoiceServices(c * gin.Context){
	log.Debug("enter func GetVoiceServices")
	appG:=  app.Gin{c}

	name :=  c.Query("name")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	valid := validation.Validation{}

	valid.MaxSize(name,100,"name").Message("名称最大不能超过100")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	voice_svc :=  voicesvc_service.VoiceService{
		Name:name,
		State:state,
		PageNum:util.GetPage(c),
		PageSize:setting.AppSetting.PageSize,
	}

	total,err := voice_svc.Count()
	if err !=  nil{
		log.Error("GetVoiceServices get svc-list by name:",name,"state:",state,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_COUNT_VOICESVC_FAIL,nil)
		return
	}

	svcs,err :=  voice_svc.GetAll()
	if err !=  nil{
		log.Error("GetVoiceServices get svc-list by name:",name,"state:",state,"err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_GET_VOICESVC_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,map[string]interface{}{
		"lists":svcs,
		"total":total,
	})
}

type AddServiceForm struct {
	Name      	string `form:"name" valid:"Required;MaxSize(100)"`
	AppId      	string `form:"appid" valid:"Required;MaxSize(100)"`
	AppKey     string `form:"appkey" valid:"Required;MaxSize(100)"`
	UserId      string `form:"userid" valid:"Required;MaxSize(100)"`
	Url      	string `form:"url" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"Required;MaxSize(300)"`
	CreatedBy 	string `form:"createdby" valid:"Required;MaxSize(100)"`
	State     	int    `form:"state" valid:"Required;Range(0,1)"`
}

// @Summary 添加服务商信息
// @Produce  json
// @Param name formData string  true  "服务商名称" maxLength(100)
// @Param appid formData string true  "服务商appid" maxLength(100)
// @Param appkey formData string true  "服务商appkey" maxLength(100)
// @Param userid formData string true  "服务商userid" maxLength(100)
// @Param url formData string true  "服务商url" maxLength(100)
// @Param description formData string true  "服务商描述信息" maxLength(300)
// @Param createdby formData string true  "本次修改人" maxLength(100)
// @Param state formData int true  "状态" Enums(0,1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/svc  [post]
func AddVoiceService(c * gin.Context){
	var (
		appG = app.Gin{C: c}
		form AddServiceForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Error("AddVoiceService  BindAndValid  err:",errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	voice_svc := voicesvc_service.VoiceService{
		Name:form.Name,
		AppId:form.AppId,
		AppKey:form.AppKey,
		UserId:form.UserId,
		URL:form.Url,
		Description:form.Description,
		CreatedBy:form.CreatedBy,
		State:form.State,
	}

	exist,err:=  voice_svc.ExistByName()
	if err !=  nil{
		log.Error("AddVoiceService  ExistByName  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_VOICESVCFAIL,nil)
		return
	}
	if exist{
		log.Error("AddVoiceService   svc   exist:",voice_svc.Name)
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_VOICESVC,nil)
		return
	}

	err=   voice_svc.Add()
	if  err !=  nil{
		log.Error("AddVoiceService    Add  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_ADD_VOICESVC_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,nil)
}

type EditServiceForm struct {
	ID          int    `form:"id" valid:"Required;Min(1)"`
	Name      	string `form:"name" valid:"MaxSize(100)"`
	AppId      	string `form:"appid" valid:"MaxSize(100)"`
	ApppKey     string `form:"appkey" valid:"MaxSize(100)"`
	UserId      string `form:"userid" valid:"MaxSize(100)"`
	Url      	string `form:"url" valid:"MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(300)"`
	ModifiedBy string  `form:"modifiedby" valid:"Required;MaxSize(100)"`
	State     	int    `form:"state" valid:"Range(0,1)"`
}
// @Summary 修改服务商信息
// @Produce  json
// @Param id path int  true  "服务商ID" mininum(1)
// @Param name formData string  false  "服务商名称" maxLength(100)
// @Param appid formData string false  "服务商appid" maxLength(100)
// @Param appkey formData string false  "服务商appkey" maxLength(100)
// @Param userid formData string false  "服务商userid" maxLength(100)
// @Param url formData string false  "服务商url" maxLength(100)
// @Param description formData string false  "服务商描述信息" maxLength(300)
// @Param modifiedby formData string true  "本次修改人" maxLength(100)
// @Param state formData int false  "状态" Enums(0,1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/svc/{id} [put]
func UpdateVoiceService(c * gin.Context){
	log.Debug("enter func UpdateVoiceService")

	var (
		appG = app.Gin{C: c}
		form = EditServiceForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Error("UpdateVoiceService    BindAndValid  errcode:",errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	voice_svc := voicesvc_service.VoiceService{
		ID:form.ID,
		Name:form.Name,
		AppId:form.AppId,
		AppKey:form.ApppKey,
		UserId:form.UserId,
		URL:form.Url,
		Description:form.Description,
		ModifiedBy:form.ModifiedBy,
		State:form.State,
	}

	exist,err:=  voice_svc.ExistByID()
	if err !=  nil{
		log.Error("UpdateVoiceService    ExistByID  errcode:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_VOICESVCFAIL,nil)
		return
	}
	if !exist{
		log.Error("UpdateVoiceService     svc  not  exist:",voice_svc.ID)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_VOICESVC,nil)
		return
	}

	err=   voice_svc.Edit()
	if  err !=  nil{
		log.Error("UpdateVoiceService   Edit  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EDIT_VOICESVC_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,nil)
}


// @Summary 删除服务商
// @Produce  json
// @Param id path int  true  "服务商ID" mininum(1)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/svc/{id} [delete]
func DelVoiceService(c* gin.Context){
	log.Debug("enter func DelVoiceService")
	appG:= app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if  valid.HasErrors(){
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
		return
	}
	voice_svc :=  voicesvc_service.VoiceService{
		ID:id,
	}

	exit,err  :=  voice_svc.ExistByID()
	if  err!=  nil{
		log.Error("DelVoiceService   ExistByID  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_VOICESVCFAIL,nil)
		return
	}
	if  !exit{
		log.Error("DelVoiceService    svc  not exist:",voice_svc.ID)
		appG.Response(http.StatusInternalServerError,e.ERROR_NOT_EXIST_VOICESVC,nil)
		return
	}

	err =voice_svc.Delete()
	if  err!=  nil{
		log.Error("DelVoiceService    Delete  err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_VOICESVC_FAIL,nil)
	}

	appG.Response(http.StatusOK,e.SUCCESS,nil)
}
