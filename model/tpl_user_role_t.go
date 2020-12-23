package model

import (
	"time"
)

type TplUserRoleT struct {
	UserId    int64     `xorm:"not null pk BIGINT(20)"`
	RoleId    int64     `xorm:"not null pk comment('功能权限') BIGINT(20)"`
	DataId    int64     `xorm:"not null default 1 comment('数据权限(默认为1所有权限)') BIGINT(20)"`
	StartTime time.Time `xorm:"comment('有效起始时间') DATETIME"`
	EndTime   time.Time `xorm:"comment('有效截止时间') DATETIME"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplUserRoleT *TplUserRoleT) tplUserRoleTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"user_id":    tplUserRoleT.UserId,
		"role_id":    tplUserRoleT.RoleId,
		"data_id":    tplUserRoleT.DataId,
		"start_time": tplUserRoleT.StartTime,
		"end_time":   tplUserRoleT.EndTime,
	}
	return respInfo
}
