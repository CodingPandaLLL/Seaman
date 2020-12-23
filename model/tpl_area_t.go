package model

import (
	"Seaman/utils"
	"time"
)

type TplAreaT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	Code             string    `xorm:"not null comment('编号') VARCHAR(100)"`
	Desp             string    `xorm:"not null comment('描述') VARCHAR(100)"`
	Type             string    `xorm:"comment('类型') VARCHAR(50)"`
	ParentId         int64     `xorm:"not null comment('父节点ID') index BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplAreaT *TplAreaT) tplAreaTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplAreaT.Id,
		"code":                tplAreaT.Code,
		"desp":                tplAreaT.Desp,
		"type":                tplAreaT.Type,
		"parent_id":           tplAreaT.ParentId,
		"revision":            tplAreaT.Revision,
		"tenant_id":           tplAreaT.TenantId,
		"app_name":            tplAreaT.AppName,
		"app_scope":           tplAreaT.AppScope,
		"create_date":         utils.FormatDatetime(tplAreaT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplAreaT.LastUpdateDate),
		"create_user_id":      tplAreaT.CreateUserId,
		"last_update_user_id": tplAreaT.LastUpdateUserId,
	}
	return respInfo
}
