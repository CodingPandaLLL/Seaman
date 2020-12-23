package model

import (
	"Seaman/utils"
	"time"
)

type TplExportHistoryT struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	TypeId           int       `xorm:"not null comment('导出类型ID') INT(11)"`
	Status           int       `xorm:"not null comment('状态') TINYINT(3)"`
	Params           string    `xorm:"comment('参数列表，使用json格式,如{"a":"b"}') VARCHAR(512)"`
	FilePath         string    `xorm:"comment('要导出的文件地址') VARCHAR(512)"`
	CreateUserId     int64     `xorm:"not null comment('创建用户ID') index BIGINT(20)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId int64     `xorm:"comment('最后更新用户ID') BIGINT(20)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	Revision         int64     `xorm:"comment('版本号') BIGINT(20)"`
	AppName          string    `xorm:"comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
	FileId           int       `xorm:"comment('文件路径ID') INT(11)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplExportHistoryT *TplExportHistoryT) tplExportHistoryTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplExportHistoryT.Id,
		"type_id":             tplExportHistoryT.TypeId,
		"status":              tplExportHistoryT.Status,
		"params":              tplExportHistoryT.Params,
		"file_path":           tplExportHistoryT.FilePath,
		"file_id":             tplExportHistoryT.FileId,
		"revision":            tplExportHistoryT.Revision,
		"tenant_id":           tplExportHistoryT.TenantId,
		"app_name":            tplExportHistoryT.AppName,
		"app_scope":           tplExportHistoryT.AppScope,
		"create_date":         utils.FormatDatetime(tplExportHistoryT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplExportHistoryT.LastUpdateDate),
		"create_user_id":      tplExportHistoryT.CreateUserId,
		"last_update_user_id": tplExportHistoryT.LastUpdateUserId,
	}
	return respInfo
}
