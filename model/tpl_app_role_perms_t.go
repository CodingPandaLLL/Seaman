package model

type TplAppRolePermsT struct {
	RoleId int64  `xorm:"not null pk comment('角色ID') BIGINT(20)"`
	Code   string `xorm:"not null pk comment('权限码') VARCHAR(20)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplAppRolePermsT *TplAppRolePermsT) tplAppRolePermsTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":   tplAppRolePermsT.RoleId,
		"code": tplAppRolePermsT.Code,
	}
	return respInfo
}
