package model

type TplUserGroupT struct {
	UserId  int64 `xorm:"not null BIGINT(20)"`
	GroupId int64 `xorm:"not null BIGINT(20)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplUserGroupT *TplUserGroupT) UserGroupTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"userId":  tplUserGroupT.UserId,
		"groupId": tplUserGroupT.GroupId,
	}
	return respInfo
}
