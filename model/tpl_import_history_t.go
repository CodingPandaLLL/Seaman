package model

import (
	"Seaman/utils"
	"time"
)

type TplImportHistoryT struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	TypeId           int       `xorm:"not null comment('导入类型ID') INT(11)"`
	Status           int       `xorm:"not null comment('状态') TINYINT(3)"`
	Params           string    `xorm:"comment('参数列表，使用json格式,如{"a":"b"}') VARCHAR(512)"`
	FilePath         string    `xorm:"comment('要导入的文件地址') VARCHAR(512)"`
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
func (tplImportHistoryT *TplImportHistoryT) tplExportTypeTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplImportHistoryT.Id,
		"type_id":             tplImportHistoryT.TypeId,
		"status":              tplImportHistoryT.Status,
		"params":              tplImportHistoryT.Params,
		"file_path":           tplImportHistoryT.FilePath,
		"file_id":             tplImportHistoryT.FileId,
		"revision":            tplImportHistoryT.Revision,
		"tenant_id":           tplImportHistoryT.TenantId,
		"app_name":            tplImportHistoryT.AppName,
		"app_scope":           tplImportHistoryT.AppScope,
		"create_date":         utils.FormatDatetime(tplImportHistoryT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplImportHistoryT.LastUpdateDate),
		"create_user_id":      tplImportHistoryT.CreateUserId,
		"last_update_user_id": tplImportHistoryT.LastUpdateUserId,
	}
	return respInfo
}
