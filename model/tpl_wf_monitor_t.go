package model

type TplWfMonitorT struct {
	Id          int64  `xorm:"pk autoincr BIGINT(20)"`
	ProcinstId  int64  `xorm:"not null comment('流程实例ID') BIGINT(20)"`
	ExecutionId int64  `xorm:"not null comment('执行ID') BIGINT(20)"`
	TaskId      string `xorm:"comment('任务ID') VARCHAR(200)"`
	TaskName    string `xorm:"comment('任务名称') VARCHAR(200)"`
	EventType   string `xorm:"not null comment('事件类型') VARCHAR(50)"`
	EventData   string `xorm:"comment('事件页面') VARCHAR(1000)"`
	Status      int    `xorm:"not null default 0 comment('状态') INT(11)"`
	HandSuggest string `xorm:"comment('处理意见') VARCHAR(1000)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplWfMonitorT *TplWfMonitorT) tplWfMonitorTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           tplWfMonitorT.Id,
		"procinst_id":  tplWfMonitorT.ProcinstId,
		"task_id":      tplWfMonitorT.TaskId,
		"task_name":    tplWfMonitorT.TaskName,
		"event_type":   tplWfMonitorT.EventType,
		"event_data":   tplWfMonitorT.EventData,
		"status":       tplWfMonitorT.Status,
		"hand_suggest": tplWfMonitorT.HandSuggest,
	}
	return respInfo
}
