package game_service

import (
	"crypto/sha256"
	"fmt"
	"gvp/models"
	log "gvp/pkg/logging"
	"math/rand"
	"time"
)

type GameService struct {
	ID   		int
	Game 		string
	GameType 	int
	AppId 		string
	AppKey 		string
	TelNum 		string
	CreatedBy 	string
	ModifiedBy 	string
	State int

	PageNum  int
	PageSize int
}




func (game *GameService) ExistByName() (bool, error) {
	return  models.ExistGameByName(game.Game)
}

func (game *GameService) ExistByID() (bool, error) {
	return  models.ExistGameByID(game.ID)
}

func (game  *  GameService) Get() (*models.Game, error) {

	vs, err := models.GetGame(game.ID)
	if err != nil {
		return nil, err
	}
	return  vs,nil
}


func (game *GameService) Add() error {
	data := make(map[string]interface{})

	appid,appkey:= game.genAppIdKey()
	log.Info("Add  game:",game.Game,"appid:",appid,"appkey:",appkey)
	data["game"] = game.Game
	data["gametype"] = game.GameType
	data["appid"] =  appid
	data["appkey"] = appkey
	data["telnum"] = game.TelNum
	data["createdby"] =  game.CreatedBy
	data["state"] = game.State

	return  models.AddGame(data)
}

func (game *GameService) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = game.ModifiedBy
	if game.State >= 0 {
		data["state"] = game.State
	}
	if game.Game != "" {
		data["game"] = game.Game
	}
	if game.TelNum != ""{
		data["telnum"] =  game.TelNum
	}
	return  models.EditGame(game.ID,data)
}

func (game *GameService) Delete() error {
	//TODO 软删除绑定关系
	return  models.DeleteGame(game.ID)
}

func (game *GameService) Count() (int, error) {
	return models.GetGameTotal(game.getMaps())
}

func (game *GameService) GetAll() ([]*models.Game, error) {
	svcs, err := models.GetGames(game.PageNum,game.PageSize,game.getMaps())
	if err != nil {
		return nil, err
	}
	return  svcs,nil
}


func (game * GameService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if game.Game != "" {
		maps["game"] = game.Game
	}
	if game.State >= 0 {
		maps["state"] = game.State
	}
	return maps
}

func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
}
func (game * GameService)genAppIdKey()(string,string){

	//TODO 生成appid
	appid := CreateCaptcha()
	//TODO 生成appkey
	str := appid +"gvp"
	sum := sha256.Sum256([]byte(str))
	appkey:= fmt.Sprintf("%x",sum)
	return   appid,appkey
}