package model

type LoginUser struct {
	UserId int64
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (loginUser *LoginUser) LoginUserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"userId": loginUser.UserId,
	}
	return respInfo
}
