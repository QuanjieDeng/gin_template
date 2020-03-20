package models

import "github.com/jinzhu/gorm"


type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool,error){
	var auth Auth
	err:=db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth).Error
	//在gorm中找不到也是中错误，所以要把这个case规避掉
	if  err !=  nil &&  err != gorm.ErrRecordNotFound {
		return  false,err
	}
	if auth.ID > 0 {
		return true,nil
	}

	return false,nil
}

