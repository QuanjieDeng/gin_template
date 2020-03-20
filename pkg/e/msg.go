package e






var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_GAME : "游戏名已经存在",
	ERROR_EXIST_GAME_FAIL: "检查游戏存在时报错",
	ERROR_NOT_EXIST_GAME : "游戏不存在",
	ERROR_GET_GAME_FAIL  : "获取游戏失败",
	ERROR_COUNT_GAME_FAIL: "获取游戏总数失败",
	ERROR_ADD_GAME_FAIL   : "增加游戏失败",
	ERROR_EDIT_GAME_FAIL   : "编辑游戏失败",
	ERROR_DELETE_GAME_FAIL : "删除游戏失败",


	ERROR_EXIST_VOICESVC       	:"服务商已经存在",
	ERROR_EXIST_VOICESVCFAIL  	:"检查服务商存在报错",
	ERROR_NOT_EXIST_VOICESVC   	:"服务商不存在",
	ERROR_GET_VOICESVC_FAIL   	: "获取服务商失败",
	ERROR_COUNT_VOICESVC_FAIL  	: "获取服务商总数失败",
	ERROR_ADD_VOICESVC_FAIL    	: "新增服务商失败",
	ERROR_EDIT_VOICESVC_FAIL   	: "编辑服务商信息失败",
	ERROR_DELETE_VOICESVC_FAIL 	: "删除服务商失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}