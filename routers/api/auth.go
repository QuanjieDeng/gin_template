package api

import (
	"net/http"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gvp/pkg/e"
	"gvp/pkg/util"
	"gvp/pkg/app"
	"gvp/pkg/setting"
	"gvp/service/auth_service"
	log  "gvp/pkg/logging"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 根据用户名和密码获取token
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} app.Response "成功"
// @Failure 500 {object} app.Response "失败"
// @Router /auth  [get]
func GetAuth(c *gin.Context) {
	appG :=   app.Gin{c}
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if   !ok{
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
		return
	}
	authService :=  auth_service.Auth{Username:username,Password:password}

	isExist,err := authService.Check()
	if    err  != nil{
		log.Error("auth-check err:",err.Error())
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH,nil)
		return
	}
	if !isExist{
		log.Error("auth-check  username:",username,"not  exist")
		appG.Response(http.StatusUnauthorized,e.ERROR_AUTH,nil)
		return
	}
	token, err := util.GenerateToken(username, password,setting.AppSetting.AuthExpire)
	if err != nil {
		log.Error("auth-check  username:",username,"GenerateToken err:",err)
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_TOKEN,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,map[string]string{
		"token":token,
	})
}