package model

import (
	"Seaman/utils"
	"time"
)

type TplFileT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	BatchNo          string    `xorm:"comment('文件标识号码') VARCHAR(36)"`
	FilePath         string    `xorm:"not null comment('文件存储相对路径') VARCHAR(255)"`
	FileName         string    `xorm:"not null comment('文件名称') VARCHAR(255)"`
	FileSize         int64     `xorm:"not null comment('文件大小') BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(32)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(32)"`
	FileUid          string    `xorm:"comment('文件UID') VARCHAR(32)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	AttachType       string    `xorm:"comment('附件类型') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplFileT *TplFileT) FileTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplFileT.Id,
		"batchNo":          tplFileT.BatchNo,
		"filePath":         tplFileT.FilePath,
		"fileName":         tplFileT.FileName,
		"fileSize":         tplFileT.FileSize,
		"fileUid":          tplFileT.FileUid,
		"attachType":       tplFileT.AttachType,
		"revision":         tplFileT.Revision,
		"tenantId":         tplFileT.TenantId,
		"appName":          tplFileT.AppName,
		"appScope":         tplFileT.AppScope,
		"createDate":       utils.FormatDatetime(tplFileT.CreateDate),
		"lastUpdateDate":   utils.FormatDatetime(tplFileT.LastUpdateDate),
		"createUserId":     tplFileT.CreateUserId,
		"lastUpdateUserId": tplFileT.LastUpdateUserId,
	}
	return respInfo
}
