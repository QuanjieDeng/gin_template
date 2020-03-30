package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GameService struct {
	ID int `gorm:"primary_key" json:"id"`
	Game string `json:"game"`
	ServiceId int`json:"serviceid" `
	GameId int `json:"gameid"`
	DeletedOn int `json:"deleted_on"`
}

//根据 服务商ID 和游戏的gameid 判断是否存在绑定关系
func ExistBindByID(svcid int,gameid int) (bool,error) {
	var gameService GameService

	err:=db.Select("id").Where("service_id = ? AND game_id = ? AND  deleted_on = ?", svcid, gameid,0).First(&gameService).Error
	if err  !=  nil &&  err != gorm.ErrRecordNotFound {
		return  false,err
	}

	if gameService.ID > 0 {
		return true,nil
	}
	return false,nil
}


func ExistBindByDBID(id int)(bool,error) {
	var gamesvc GameService
	err:=db.Select("id").Where("id = ?", id).First(&gamesvc).Error
	if  err  !=  nil &&  err != gorm.ErrRecordNotFound {
		return   false,err
	}

	if gamesvc.ID > 0 {
		return true,nil
	}

	return false,nil
}


//获取绑定关系的总数
func GetGameServiceTotal(maps interface {}) (int,error){
	var  count  int
	err:=db.Model(&GameService{}).Where(maps).Count(&count).Error
	if  err !=  nil{
		return 0,err
	}
	return  count,nil
}

//获取绑定关系列表
func GetGameServices(pageNum int, pageSize int, maps interface{}) ([]*GameService, error) {
	var (
		gameservice []*GameService
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&gameservice).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&gameservice).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return gameservice, nil
}



//根据ID获取绑定关系数据
func GetGameService(id int) (* GameService, error) {
	var  gameservice  GameService
	err:= db.Where("id = ?", id).First(&gameservice).Error
	if  err != nil{
		return nil,err
	}

	return  &gameservice,nil
}



//增加一个绑定关系
func AddGameService(data map[string]interface {})  error {
	err:= db.Create(&GameService {
		Game : data["game"].(string),
		ServiceId : data["serviceid"].(int),
		GameId : data["gameid"].(int),
	}).Error

	if err != nil{
		return  err
	}
	return  nil
}

//软删除-根据ID
func DeleteGameService(id int) error {
	data := make(map[string]interface{})
	data["deleted_on"] =   time.Now().Unix()
	err:= db.Model(&GameService{}).Where("id = ?", id).Updates(data).Error
	if  err !=  nil{
		return  err
	}
	return nil
}

//软删除-根据VoiceService ID
func DeleteGameServiceBySvcId(id int) error {
	data := make(map[string]interface{})
	data["deleted_on"] =   time.Now().Unix()
	err:= db.Model(&GameService{}).Where("service_id = ?", id).Updates(data).Error
	if  err !=  nil{
		return  err
	}
	return nil
}

//软删除-根据game ID
func DeleteGameServiceByGameId(id int) error {
	data := make(map[string]interface{})
	data["deleted_on"] =   time.Now().Unix()
	err:= db.Model(&GameService{}).Where("game_id = ?", id).Updates(data).Error
	if  err !=  nil{
		return  err
	}
	return nil
}



//硬删除
func DeleteGameServiceHard(id int)   error {
	err:= db.Where("id = ?", id).Delete(GameService{}).Error
	if  err  !=  nil{
		return  err
	}
	return   nil
}


//根据游戏开发商的appid,查询他们关联的服务商的信息

func GetSvcByGameAppId(appid string,svcname  string)(*VoiceService,error){

	var  svc  VoiceService
	err := db.Raw("SELECT a.* from gvp_voice_service a ,gvp_game b,gvp_game_service c" +
		"  where a.id=c.service_id and b.id=c.game_id and a.state=1 and b.app_id=? and a.name=?",appid,svcname).Scan(&svc).Error
	if err !=nil{
		return  nil,err
	}
	return  &svc,nil
}