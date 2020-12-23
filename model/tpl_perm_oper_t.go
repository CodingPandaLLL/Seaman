package model

import (
	"Seaman/utils"
	"time"
)

type TplPermOperT struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Code           string    `xorm:"not null comment('操作码编号') VARCHAR(16)"`
	Status         int       `xorm:"not null default 1 comment('状态(0:无效,1:有效)') INT(11)"`
	Desp           string    `xorm:"comment('描述') VARCHAR(500)"`
	ParentId       int64     `xorm:"not null comment('资源ID') BIGINT(20)"`
	Revision       int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateDate     time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId       string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName        string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope       string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplPermOperT *TplPermOperT) tplPermOperTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplPermOperT.Id,
		"code":             tplPermOperT.Code,
		"status":           tplPermOperT.Status,
		"desp":             tplPermOperT.Desp,
		"parent_id":        tplPermOperT.ParentId,
		"revision":         tplPermOperT.Revision,
		"tenant_id":        tplPermOperT.TenantId,
		"app_name":         tplPermOperT.AppName,
		"app_scope":        tplPermOperT.AppScope,
		"create_date":      utils.FormatDatetime(tplPermOperT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplPermOperT.LastUpdateDate),
	}
	return respInfo
}
