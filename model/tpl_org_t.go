package model

import (
	"Seaman/utils"
	"time"
)

type TplOrgT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	PId              int64     `xorm:"comment('父ID') BIGINT(20)"`
	OrgName          string    `xorm:"not null comment('部门名称') VARCHAR(255)"`
	OrgCode          string    `xorm:"not null comment('部门编码') VARCHAR(255)"`
	OrgNote          string    `xorm:"comment('部门节点') VARCHAR(255)"`
	Filed1           string    `xorm:"VARCHAR(255)"`
	Filed2           time.Time `xorm:"DATE"`
	Filed3           int64     `xorm:"BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建人ID') VARCHAR(32)"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新人ID') VARCHAR(32)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplOrgT *TplOrgT) tplOrgTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplOrgT.Id,
		"p_id":                tplOrgT.PId,
		"org_name":            tplOrgT.OrgName,
		"org_code":            tplOrgT.OrgCode,
		"filed1":              tplOrgT.Filed1,
		"filed2":              tplOrgT.Filed2,
		"filed3":              tplOrgT.Filed3,
		"tenant_id":           tplOrgT.TenantId,
		"app_name":            tplOrgT.AppName,
		"app_scope":           tplOrgT.AppScope,
		"create_date":         utils.FormatDatetime(tplOrgT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplOrgT.LastUpdateDate),
		"create_user_id":      tplOrgT.CreateUserId,
		"last_update_user_id": tplOrgT.LastUpdateUserId,
	}
	return respInfo
}
