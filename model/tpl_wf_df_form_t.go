package model

import (
	"Seaman/utils"
	"time"
)

type TplWfDfFormT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	JsonData         string    `xorm:"comment('json数据') VARCHAR(1000)"`
	FormKey          string    `xorm:"comment('表单页面地址') VARCHAR(1000)"`
	ProcDefId        string    `xorm:"not null comment('流程定义ID') VARCHAR(64)"`
	ProcDefName      string    `xorm:"not null comment('流程定义名称') VARCHAR(255)"`
	ProcDefVersion   int       `xorm:"not null comment('流程定义版本') INT(11)"`
	Category         string    `xorm:"VARCHAR(50)"`
	Name             string    `xorm:"comment('流程名称') VARCHAR(64)"`
	ProcDefKey       string    `xorm:"comment('流程定义KEY') VARCHAR(64)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('流程创建人ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新人ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplWfDfFormT *TplWfDfFormT) tplWfDfFormTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplWfDfFormT.Id,
		"json_data":        tplWfDfFormT.JsonData,
		"form_key":         tplWfDfFormT.FormKey,
		"proc_def_id":      tplWfDfFormT.ProcDefId,
		"proc_def_name":    tplWfDfFormT.ProcDefName,
		"proc_def_version": tplWfDfFormT.ProcDefVersion,
		"category":         tplWfDfFormT.Category,
		"name":             tplWfDfFormT.Name,
		"proc_def_key":     tplWfDfFormT.ProcDefKey,
		"revision":         tplWfDfFormT.Revision,
		"tenant_id":        tplWfDfFormT.TenantId,
		"app_name":         tplWfDfFormT.AppName,
		"create_date":      utils.FormatDatetime(tplWfDfFormT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplWfDfFormT.LastUpdateDate),
		"create_user_id":   tplWfDfFormT.LastUpdateUserId,
	}
	return respInfo
}
