package bind_service

import (
	"gvp/models"
)

type BindService struct {
	ID   		int
	Game 		string
	ServiceId 	int
	GameId 		int

	PageNum  int
	PageSize int
}




func (bind *BindService) ExistBySerIdAndGameId() (bool, error) {
	return  models.ExistBindByID(bind.ServiceId,bind.GameId)
}


func (bind *BindService) ExistByID() (bool, error) {
	return  models.ExistBindByDBID(bind.ID)
}


func (bind  *  BindService) Get() (*models.GameService, error) {

	vs, err := models.GetGameService(bind.ID)
	if err != nil {
		return nil, err
	}
	return  vs,nil
}


func (bind *BindService) Add() error {
	data := make(map[string]interface{})

	data["game"] 		= bind.Game
	data["serviceid"] 	= bind.ServiceId
	data["gameid"] 		= bind.GameId

	return  models.AddGameService(data)
}


func (bind *BindService) Delete() error {
	return  models.DeleteGameService(bind.ID)
}

func (bind *BindService) Count() (int, error) {
	return models.GetGameServiceTotal(bind.getMaps())
}

func (bind *BindService) GetAll() ([]*models.GameService, error) {
	svcs, err := models.GetGameServices(bind.PageNum,bind.PageSize,bind.getMaps())
	if err != nil {
		return nil, err
	}
	return  svcs,nil
}


func (bind * BindService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if bind.Game != "" {
		maps["game"] = bind.Game
	}

	if bind.ServiceId >= 0 {
		maps["service_id"] =   bind.ServiceId
	}
	if bind.GameId >= 0 {
		maps["game_id"] =   bind.GameId
	}

	return maps
}
