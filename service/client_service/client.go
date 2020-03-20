package client_service

import (
	"gvp/models"
)

type  Client struct {

	AppID 	string
	AppKey 	string
	VoiceSvcName string
}

//判断用户是否存在
func (cli * Client)Exist()(bool,error){
	return models.ExistGameAppIdKey(cli.AppID,cli.AppKey)
	return true,nil
}


//校验用户身份
func(cli * Client)Check()(bool,error){

	return true,nil
}


//获取客户端对应的VoiceServer信息
func  (cli * Client)GetVoiceService()(* models.VoiceService,error){
	return  models.GetSvcByGameAppId(cli.AppID,cli.VoiceSvcName)
}