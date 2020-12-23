package model

import (
	"Seaman/utils"
	"time"
)

type TplImportTypeT struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	ImportType       string    `xorm:"not null comment('类型') unique VARCHAR(64)"`
	Description      string    `xorm:"not null comment('描述') VARCHAR(256)"`
	TemplateFilePath string    `xorm:"comment('模板文件') VARCHAR(512)"`
	Sync             int       `xorm:"not null comment('是否同步(0:异步,1:同步)') TINYINT(1)"`
	ExeClass         string    `xorm:"not null comment('执行导入操作的类') VARCHAR(128)"`
	CreateUserId     int64     `xorm:"not null comment('创建用户ID') BIGINT(20)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	Pager            int       `xorm:"comment('是否分页导入(0:不分页,1:分页)') TINYINT(1)"`
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
func (tplImportTypeT *TplImportTypeT) tplImportTypeTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplImportTypeT.Id,
		"import_type":         tplImportTypeT.ImportType,
		"description":         tplImportTypeT.Description,
		"template_file_path":  tplImportTypeT.TemplateFilePath,
		"sync":                tplImportTypeT.Sync,
		"exec_lass":           tplImportTypeT.ExeClass,
		"pager":               tplImportTypeT.Pager,
		"size":                tplImportTypeT.Size,
		"file_id":             tplImportTypeT.FileId,
		"revision":            tplImportTypeT.Revision,
		"tenant_id":           tplImportTypeT.TenantId,
		"app_name":            tplImportTypeT.AppName,
		"app_scope":           tplImportTypeT.AppScope,
		"create_date":         utils.FormatDatetime(tplImportTypeT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplImportTypeT.LastUpdateDate),
		"create_user_id":      tplImportTypeT.CreateUserId,
		"last_update_user_id": tplImportTypeT.LastUpdateUserId,
	}
	return respInfo
}
