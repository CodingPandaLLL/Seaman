package model

type TplUserGroupT struct {
	UserId  int64 `xorm:"not null BIGINT(20)"`
	GroupId int64 `xorm:"not null BIGINT(20)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplUserGroupT *TplUserGroupT) tplUserGroupTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"user_id":  tplUserGroupT.UserId,
		"group_id": tplUserGroupT.GroupId,
	}
	return respInfo
}
