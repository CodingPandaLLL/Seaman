package model

import (
	"Seaman/utils"
	"time"
)

type TplDpResourceT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	ResourceDesc     string    `xorm:"not null comment('资源描述') VARCHAR(200)"`
	ResourceName     string    `xorm:"not null comment('资源物理名称') VARCHAR(200)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplDpResourceT *TplDpResourceT) tplDpResourceTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplDpResourceT.Id,
		"resource_desc":       tplDpResourceT.ResourceDesc,
		"resource_name":       tplDpResourceT.ResourceName,
		"revision":            tplDpResourceT.Revision,
		"tenant_id":           tplDpResourceT.TenantId,
		"app_name":            tplDpResourceT.AppName,
		"create_date":         utils.FormatDatetime(tplDpResourceT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplDpResourceT.LastUpdateDate),
		"create_user_id":      tplDpResourceT.CreateUserId,
		"last_update_user_id": tplDpResourceT.LastUpdateUserId,
	}
	return respInfo
}
