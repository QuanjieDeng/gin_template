package models

import "github.com/jinzhu/gorm"

type Game struct {
	Model
	Game string `json:"game"`
	GameType int `json:"gametype"`
	AppId string `json:"appid"`
	AppKey string `json:"appkey"`
	TelNum string `json:"telnum"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}


//根据name判断 游戏是否存在
func ExistGameByName(name string)(bool,error) {
	var game  Game
	err:=db.Select("id").Where("game = ?", name).First(&game).Error
	if  err  !=  nil &&  err != gorm.ErrRecordNotFound {
		return   false,err
	}

	if game.ID > 0 {
		return true,nil
	}
	return false,nil
}

//根据ID判断 游戏是否存在
func ExistGameByID(id int)(bool,error) {
	var game  Game
	err:=db.Select("id").Where("id = ?", id).First(&game).Error
	if  err  !=  nil &&  err != gorm.ErrRecordNotFound {
		return   false,err
	}

	if game.ID > 0 {
		return true,nil
	}

	return false,nil
}

//根据appid appkey 判断 游戏是否存在
func ExistGameAppIdKey(appid,appkey string)(bool,error) {
	var game  Game
	err:=db.Select("id").Where("app_id = ? AND app_key = ?",appid,appkey).First(&game).Error
	if  err  !=  nil &&  err != gorm.ErrRecordNotFound {
		return   false,err
	}

	if game.ID > 0 {
		return true,nil
	}

	return false,nil
}



//获取游戏开发商总数
func GetGameTotal(maps interface {}) (int,error){
	var   count  int
	err:=db.Model(&Game{}).Where(maps).Count(&count).Error
	if  err != nil{
		return  0,err
	}
	return  count,nil
}

//根据分页要求获取游戏开发商列表
func GetGames(pageNum int, pageSize int, maps interface {}) ([]*Game,error) {
	var  games  []*Game
	var err  error
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&games).Error
	} else {
		err = db.Where(maps).Find(&games).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return  games,nil
}

//根据ID获取开发生对象
func GetGame(id int) (*Game,error) {
	var  game  Game
	err:= db.Where("id = ?", id).First(&game).Error
	if  err !=  nil{
		return  nil,err
	}

	return  &game,nil
}

//更新游戏开放商数据
func EditGame(id int, data interface {}) error {
	err:= db.Model(&Game{}).Where("id = ?", id).Updates(data).Error
	if  err !=  nil{
		return  err
	}
	return  nil
}


//新增游戏开发商
func AddGame(data map[string]interface {}) error {
	err:=db.Create(&Game {
		Game : data["game"].(string),
		GameType : data["gametype"].(int),
		AppId : data["appid"].(string),
		AppKey : data["appkey"].(string),
		TelNum : data["telnum"].(string),
		CreatedBy : data["createdby"].(string),
		State : data["state"].(int),
	}).Error

	if  err !=  nil{
		return  err
	}

	return   nil
}


//删除一个游戏开发商
func DeleteGame(id int) error {
	err:= db.Where("id = ?", id).Delete(Game{}).Error
	if  err !=  nil{
		return  err
	}
	return nil
}
