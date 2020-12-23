package model

import (
	"Seaman/utils"
	"time"
)

type TplWfRuFormT struct {
	Id               string    `xorm:"not null pk VARCHAR(64)"`
	ProcInstId       string    `xorm:"not null comment('流程实例ID') unique VARCHAR(64)"`
	BusinessKey      string    `xorm:"not null comment('业务ID') VARCHAR(64)"`
	ProcDefId        string    `xorm:"not null comment('流程定义ID') VARCHAR(64)"`
	ProcDefName      string    `xorm:"not null comment('流程定义名称') VARCHAR(255)"`
	ProcDefVersion   int       `xorm:"not null comment('流程定义版本') INT(11)"`
	Category         string    `xorm:"VARCHAR(50)"`
	StartUserId      string    `xorm:"comment('启动流程用户ID') VARCHAR(36)"`
	Name             string    `xorm:"comment('流程名称') VARCHAR(255)"`
	Status           string    `xorm:"comment('状态') VARCHAR(255)"`
	Handler          string    `xorm:"comment('处理人（群组和用户）') VARCHAR(4000)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('流程创建人ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新人ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	ProcTaskSize     int       `xorm:"comment('流程实例任务数') INT(11)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplWfRuFormT *TplWfRuFormT) tplWfRuFormTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplWfRuFormT.Id,
		"ProcInstId":          tplWfRuFormT.ProcInstId,
		"BusinessKey":         tplWfRuFormT.BusinessKey,
		"proc_def_id":         tplWfRuFormT.ProcDefId,
		"proc_def_name":       tplWfRuFormT.ProcDefName,
		"proc_def_version":    tplWfRuFormT.ProcDefVersion,
		"category":            tplWfRuFormT.Category,
		"start_user_id":       tplWfRuFormT.StartUserId,
		"name":                tplWfRuFormT.Name,
		"status":              tplWfRuFormT.Status,
		"handler":             tplWfRuFormT.Handler,
		"proc_task_size":      tplWfRuFormT.ProcTaskSize,
		"tenant_id":           tplWfRuFormT.TenantId,
		"revision":            tplWfRuFormT.Revision,
		"app_name":            tplWfRuFormT.AppName,
		"app_scope":           tplWfRuFormT.AppScope,
		"create_date":         utils.FormatDatetime(tplWfRuFormT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplWfRuFormT.LastUpdateDate),
		"last_update_user_id": tplWfRuFormT.LastUpdateUserId,
		"create_user_id":      tplWfRuFormT.CreateUserId,
	}
	return respInfo
}
