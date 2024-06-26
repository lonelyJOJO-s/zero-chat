package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[REDIS_ERROR] = "redis数据库错误"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[NO_ACCESS_TO_RESOURCE] = "no access to the current resources"
	message[INSERT_ALREADY_EXSIT] = "插入对象已存在，请勿重复添加"
	message[UserNotInGroup] = "用户不在当前的群组中"
	message[USER_NOT_FOUND] = "当前用户不存在"
	message[MUST_CHOOSE_HEIR] = "必须选择一个候选人"
	message[WEBSOCKET_CONN_ERR] = "websocket链接失败"
	message[UNACCESSABLE_ERROR] = "401未授权"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
