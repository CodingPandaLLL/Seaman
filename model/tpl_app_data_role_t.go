package model

import (
	"Seaman/utils"
	"time"
)

type TplAppDataRoleT struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Name           string    `xorm:"not null comment('角色名称') VARCHAR(128)"`
	Desp           string    `xorm:"not null comment('角色描述') VARCHAR(128)"`
	Status         int64     `xorm:"not null comment('状态') BIGINT(20)"`
	Defaultin      string    `xorm:"not null default '0' comment('系统内置标识（0否1是）') CHAR(1)"`
	Code           string    `xorm:"comment('编码') VARCHAR(128)"`
	TenantId       string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName        string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope       string    `xorm:"comment('系统群名') VARCHAR(32)"`
	CreateDate     time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate time.Time `xorm:"comment('最后修改时间') DATETIME"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplAppDataRoleT *TplAppDataRoleT) tplAppDataRoleTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplAppDataRoleT.Id,
		"name":             tplAppDataRoleT.Name,
		"desp":             tplAppDataRoleT.Desp,
		"defaultin":        tplAppDataRoleT.Defaultin,
		"code":             tplAppDataRoleT.Code,
		"tenant_id":        tplAppDataRoleT.TenantId,
		"app_name":         tplAppDataRoleT.AppName,
		"app_scope":        tplAppDataRoleT.AppScope,
		"create_date":      utils.FormatDatetime(tplAppDataRoleT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplAppDataRoleT.LastUpdateDate),
	}
	return respInfo
}
