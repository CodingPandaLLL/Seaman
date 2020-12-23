package model

import (
	"time"
)

type TplWfLogT struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	ProcinstId     string    `xorm:"not null comment('流程实例ID') VARCHAR(20)"`
	TaskId         string    `xorm:"not null comment('任务ID') VARCHAR(20)"`
	TaskName       string    `xorm:"comment('任务名称') VARCHAR(255)"`
	TaskCreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	TaskEndTime    time.Time `xorm:"comment('最后修改时间') DATETIME"`
	Operation      string    `xorm:"comment('操作记录') VARCHAR(255)"`
	Operator       string    `xorm:"not null comment('任务操作人') VARCHAR(20)"`
	HandSuggest    string    `xorm:"comment('处理意见') VARCHAR(255)"`
	OperationType  string    `xorm:"comment('操作类型') VARCHAR(10)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplWfLogT *TplWfLogT) tplWfLogTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplWfLogT.Id,
		"procinst_id":      tplWfLogT.ProcinstId,
		"task_id":          tplWfLogT.TaskId,
		"task_name":        tplWfLogT.TaskName,
		"task_create_time": tplWfLogT.TaskCreateTime,
		"task_end_time":    tplWfLogT.TaskEndTime,
		"operation":        tplWfLogT.Operation,
		"operator":         tplWfLogT.Operator,
		"hand_suggest":     tplWfLogT.HandSuggest,
		"operation_type":   tplWfLogT.OperationType,
	}
	return respInfo
}
