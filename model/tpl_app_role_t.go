package model

import (
	"Seaman/utils"
	"time"
)

type TplAppRoleT struct {
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
func (tplAppRoleT *TplAppRoleT) tplAppRoleTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplAppRoleT.Id,
		"name":             tplAppRoleT.Name,
		"desp":             tplAppRoleT.Desp,
		"status":           tplAppRoleT.Status,
		"defaultin":        tplAppRoleT.Defaultin,
		"code":             tplAppRoleT.Code,
		"tenant_id":        tplAppRoleT.TenantId,
		"app_name":         tplAppRoleT.AppName,
		"app_scope":        tplAppRoleT.AppScope,
		"create_date":      utils.FormatDatetime(tplAppRoleT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplAppRoleT.LastUpdateDate),
	}
	return respInfo
}
