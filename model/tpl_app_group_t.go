package model

import (
	"Seaman/utils"
	"time"
)

type TplAppGroupT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	Code             string    `xorm:"not null comment('群组编号') VARCHAR(128)"`
	Name             string    `xorm:"not null comment('群组名称') VARCHAR(128)"`
	Status           int       `xorm:"not null comment('状态') INT(11)"`
	Defaultin        int       `xorm:"not null default 0 comment('系统内置标识（0否1是）') INT(11)"`
	Desp             string    `xorm:"comment('描述') VARCHAR(500)"`
	Type             string    `xorm:"comment('类型') VARCHAR(32)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(32)"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(32)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplAppGroupT *TplAppGroupT) tplAppGroupTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplAppGroupT.Id,
		"code":                tplAppGroupT.Code,
		"name":                tplAppGroupT.Name,
		"status":              tplAppGroupT.Status,
		"defaultin":           tplAppGroupT.Defaultin,
		"desp":                tplAppGroupT.Desp,
		"type":                tplAppGroupT.Type,
		"revision":            tplAppGroupT.Revision,
		"tenant_id":           tplAppGroupT.TenantId,
		"app_name":            tplAppGroupT.AppName,
		"app_scope":           tplAppGroupT.AppScope,
		"create_date":         utils.FormatDatetime(tplAppGroupT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplAppGroupT.LastUpdateDate),
		"create_user_id":      tplAppGroupT.CreateUserId,
		"last_update_user_id": tplAppGroupT.LastUpdateUserId,
	}
	return respInfo
}
