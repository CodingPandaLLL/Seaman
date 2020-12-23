package model

import (
	"Seaman/utils"
	"time"
)

type TplWfHiFormT struct {
	Id               string    `xorm:"not null pk VARCHAR(64)"`
	ProcInstId       string    `xorm:"not null comment('流程实例ID') unique VARCHAR(64)"`
	BusinessKey      string    `xorm:"not null comment('业务ID') VARCHAR(64)"`
	ProcDefId        string    `xorm:"not null comment('流程定义ID') VARCHAR(64)"`
	ProcDefName      string    `xorm:"not null comment('流程定义名称') VARCHAR(255)"`
	ProcDefVersion   int       `xorm:"not null comment('流程定义版本') INT(11)"`
	Category         string    `xorm:"VARCHAR(50)"`
	StartUserId      string    `xorm:"comment('流程启动用户ID') VARCHAR(36)"`
	DeleteReason     string    `xorm:"comment('删除原因') VARCHAR(4000)"`
	EndDate          time.Time `xorm:"comment('流程结束日期') DATETIME"`
	Duration         int64     `xorm:"comment('持续时间') BIGINT(20)"`
	Name             string    `xorm:"comment('名称') VARCHAR(255)"`
	Status           string    `xorm:"comment('状态') VARCHAR(255)"`
	Handler          string    `xorm:"comment('处理人（群组和用户）') VARCHAR(4000)"`
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
func (tplWfHiFormT *TplWfHiFormT) tplWfHiFormTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplWfHiFormT.Id,
		"ProcInstId":       tplWfHiFormT.ProcInstId,
		"BusinessKey":      tplWfHiFormT.BusinessKey,
		"proc_def_id":      tplWfHiFormT.ProcDefId,
		"proc_def_name":    tplWfHiFormT.ProcDefName,
		"proc_def_version": tplWfHiFormT.ProcDefVersion,
		"category":         tplWfHiFormT.Category,
		"start_user_id":    tplWfHiFormT.StartUserId,
		"delete_reason":    tplWfHiFormT.DeleteReason,
		"end_date":         tplWfHiFormT.EndDate,
		"duration":         tplWfHiFormT.Duration,
		"name":             tplWfHiFormT.Name,
		"status":           tplWfHiFormT.Status,
		"handler":          tplWfHiFormT.Handler,
		"proc_task_size":   tplWfHiFormT.ProcTaskSize,
		"tenant_id":        tplWfHiFormT.TenantId,
		"app_name":         tplWfHiFormT.AppName,
		"create_date":      utils.FormatDatetime(tplWfHiFormT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplWfHiFormT.LastUpdateDate),
		"create_user_id":   tplWfHiFormT.LastUpdateUserId,
	}
	return respInfo
}
