package model

import (
	"Seaman/utils"
	"time"
)

type TplPermResourceT struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Code           string    `xorm:"not null comment('资源编号') VARCHAR(16)"`
	Status         int       `xorm:"not null default 1 comment('状态(0:无效,1:有效)') INT(11)"`
	Desp           string    `xorm:"comment('描述') VARCHAR(500)"`
	Revision       int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateDate     time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId       string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName        string    `xorm:"not null comment('应用(模块)名') VARCHAR(32)"`
	AppScope       string    `xorm:"comment('系统群名') VARCHAR(32)"`
	AppCode        string    `xorm:"not null comment('应用(模块)编码') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplPermResourceT *TplPermResourceT) tplPermResourceTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplPermResourceT.Id,
		"code":             tplPermResourceT.Code,
		"status":           tplPermResourceT.Status,
		"desp":             tplPermResourceT.Desp,
		"revision":         tplPermResourceT.Revision,
		"tenant_id":        tplPermResourceT.TenantId,
		"app_name":         tplPermResourceT.AppName,
		"app_scope":        tplPermResourceT.AppScope,
		"create_date":      utils.FormatDatetime(tplPermResourceT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplPermResourceT.LastUpdateDate),
	}
	return respInfo
}
