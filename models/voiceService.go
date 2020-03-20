
package models

import "github.com/jinzhu/gorm"

type VoiceService struct {
	Model
	Name string `json:"name"`
	AppId string `json:"appid"`
	AppKey string `json:"appkey"`
	UserId string `json:"userid"`
	URL string `json:"url"`
	Description string `json:"description"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`

	GameService []GameService
}

//根据名字判断服务商是否存在
func ExistVoiceServiceByName(name  string) (bool,error) {
	var voiceservice VoiceService
	err:=db.Select("id").Where("name = ?", name).First(&voiceservice).Error

	if  err !=  nil &&  err != gorm.ErrRecordNotFound {
		return  false,err
	}

	if voiceservice.ID > 0 {
		return true,nil
	}

	return false,nil
}

//根据数据库ID判断服务商是否存在
func ExistVoiceServiceById(id  int)(bool,error) {
	var voiceservice VoiceService
	err:=db.Select("id").Where("id = ?",id).First(&voiceservice).Error
	if  err !=  nil &&  err != gorm.ErrRecordNotFound {
		return  false,err
	}

	if voiceservice.ID > 0 {
		return true,nil
	}

	return false,nil
}

//获取服务商的总数
func GetVoiceServiceTotal(maps interface {}) (int,error){
	var  count  int
	err:= db.Model(&VoiceService{}).Where(maps).Count(&count).Error
	if  err !=  nil{
		return  0,err
	}

	return count,nil
}

//根据ID获取单个服务商的信息
func GetVoiceServicebyId(id int) (*VoiceService, error) {
	var   svc   VoiceService
	err:= db.Where("id = ?", id).First(&svc).Error
	if  err !=  nil{
		return  nil,err
	}
	return   &svc,nil
}

//根据每页的展示数量获取服务商列表
func GetVoiceServices(pageNum int, pageSize int, maps interface {}) ([]*VoiceService,error) {
	var   svc_list []*VoiceService
	var  err  error

	if pageSize > 0 && pageNum > 0 {
		err =db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&svc_list).Error
	} else {
		err = db.Where(maps).Find(&svc_list).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return  svc_list,nil
}


//更新服务商的信息
func EditVoiceService(id int, data interface {}) (error) {

	err:= db.Model(&VoiceService{}).Where("id = ?", id).Updates(data).Error
	if  err !=  nil{
		return   err
	}

	return nil
}

//添加一个服务商
func AddVoiceService(data map[string]interface {}) error {
	err:= db.Create(&VoiceService {
		Name : data["name"].(string),
		AppId : data["appid"].(string),
		AppKey : data["appkey"].(string),
		UserId : data["userid"].(string),
		URL : data["url"].(string),
		Description : data["description"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	}).Error

	if  err  !=  nil{
		return  err
	}

	return  nil
}


//删除一个服务商
func DeleteVoiceService(id int) error {
	err  := db.Where("id = ?", id).Delete(VoiceService{}).Error
	if err != nil{
		return  err
	}

	return  nil
}
