package voicesvc_service

import (
	"gvp/models"
)


type VoiceService struct {
	ID   int
	Name string
	AppId string
	AppKey string
	UserId string
	URL string
	Description string
	CreatedBy string
	ModifiedBy string
	State int

	PageNum  int
	PageSize int
}



func (svc *VoiceService) ExistByName() (bool, error) {
	return  models.ExistVoiceServiceByName(svc.Name)
}

func (svc *VoiceService) ExistByID() (bool, error) {
	return  models.ExistVoiceServiceById(svc.ID)
}

func (svc  *  VoiceService) Get() (*models.VoiceService, error) {

	vs, err := models.GetVoiceServicebyId(svc.ID)
	if err != nil {
		return nil, err
	}
	return  vs,nil
}



func (svc *VoiceService) Add() error {
	data := make(map[string]interface{})

	data["name"] = svc.Name
	data["appid"] = svc.AppId
	data["appkey"] =   svc.AppKey
	data["userid"] = svc.UserId
	data["url"] =  svc.URL
	data["description"] =   svc.Description
	data["created_by"] =  svc.CreatedBy
	data["state"] =  svc.State
	return  models.AddVoiceService(data)
}

func (svc *VoiceService) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = svc.ModifiedBy
	if svc.State >= 0 {
		data["state"] = svc.State
	}
	if svc.Name != "" {
		data["name"] = svc.Name
	}
	if svc.AppId != ""{
		data["appid"] = svc.AppId
	}
	if svc.AppKey != ""{
		data["appkey"] = svc.AppKey
	}
	if svc.UserId != ""{
		data["userid"] =  svc.UserId
	}
	if  svc.URL != ""{
		data["url"] =  svc.URL
	}
	if svc.Description !=  ""{
		data["description"] = svc.Description
	}
	return  models.EditVoiceService(svc.ID,data)
}

func (svc *VoiceService) Delete() error {
	//软删除绑定关系
	models.DeleteGameServiceBySvcId(svc.ID)
	return  models.DeleteVoiceService(svc.ID)
}

func (svc *VoiceService) Count() (int, error) {
	return models.GetVoiceServiceTotal(svc.getMaps())
}

func (svc *VoiceService) GetAll() ([]*models.VoiceService, error) {
	svcs, err := models.GetVoiceServices(svc.PageNum, svc.PageSize, svc.getMaps())
	if err != nil {
		return nil, err
	}
	return  svcs,nil
}


func (svc * VoiceService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if svc.Name != "" {
		maps["name"] = svc.Name
	}
	if svc.State >= 0 {
		maps["state"] = svc.State
	}
	return maps
}
