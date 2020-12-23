package model

import (
	"Seaman/utils"
	"time"
)

type TplExportTypeT struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	ExportType       string    `xorm:"not null comment('类型') unique VARCHAR(64)"`
	Description      string    `xorm:"not null comment('描述') VARCHAR(256)"`
	TemplateFilePath string    `xorm:"comment('模板文件') VARCHAR(512)"`
	Sync             int       `xorm:"not null comment('是否同步(0:异步,1:同步)') TINYINT(1)"`
	ExeClass         string    `xorm:"not null comment('执行导出操作的类') VARCHAR(128)"`
	CreateUserId     int64     `xorm:"not null comment('创建用户ID') BIGINT(20)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	Pager            int       `xorm:"comment('是否分页导出(0:不分页,1:分页)') TINYINT(1)"`
	Size             int       `xorm:"comment('每页条数') INT(11)"`
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
func (tplExportTypeT *TplExportTypeT) tplExportTypeTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplExportTypeT.Id,
		"export_type":         tplExportTypeT.ExportType,
		"description":         tplExportTypeT.Description,
		"template_file_path":  tplExportTypeT.TemplateFilePath,
		"sync":                tplExportTypeT.Sync,
		"exe_class":           tplExportTypeT.ExeClass,
		"pager":               tplExportTypeT.Pager,
		"size":                tplExportTypeT.Size,
		"file_id":             tplExportTypeT.FileId,
		"revision":            tplExportTypeT.Revision,
		"tenant_id":           tplExportTypeT.TenantId,
		"app_name":            tplExportTypeT.AppName,
		"app_scope":           tplExportTypeT.AppScope,
		"create_date":         utils.FormatDatetime(tplExportTypeT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplExportTypeT.LastUpdateDate),
		"create_user_id":      tplExportTypeT.CreateUserId,
		"last_update_user_id": tplExportTypeT.LastUpdateUserId,
	}
	return respInfo
}
