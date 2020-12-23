package model

import (
	"Seaman/utils"
	"time"
)

type TplBuildTemplateLogT struct {
	Id              int       `xorm:"not null pk autoincr INT(11)"`
	ProjectName     string    `xorm:"not null comment('项目名') VARCHAR(64)"`
	DownloadUrl     string    `xorm:"comment('下载链接') VARCHAR(500)"`
	Status          string    `xorm:"comment('状态 1 有效 0 无效') VARCHAR(100)"`
	TemplateContent string    `xorm:"comment('模板内容') TEXT"`
	Description     string    `xorm:"comment('简单描述') VARCHAR(500)"`
	CreateUserId    int64     `xorm:"not null comment('创建用户ID') BIGINT(20)"`
	CreateDate      time.Time `xorm:"comment('创建时间') DATETIME"`
	Revision        int64     `xorm:"comment('版本号') BIGINT(20)"`
	AppName         string    `xorm:"comment('应用名') VARCHAR(32)"`
	TenantId        string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope        string    `xorm:"comment('系统群名') VARCHAR(32)"`
	LastUpdateDate  time.Time `xorm:"comment('最后修改时间') DATETIME"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplBuildTemplateLogT *TplBuildTemplateLogT) tplBuildTemplateLogTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               tplBuildTemplateLogT.Id,
		"Project_Name":     tplBuildTemplateLogT.ProjectName,
		"download_url":     tplBuildTemplateLogT.DownloadUrl,
		"status":           tplBuildTemplateLogT.Status,
		"template_content": tplBuildTemplateLogT.TemplateContent,
		"description":      tplBuildTemplateLogT.Description,
		"revision":         tplBuildTemplateLogT.Revision,
		"tenant_id":        tplBuildTemplateLogT.TenantId,
		"app_name":         tplBuildTemplateLogT.AppName,
		"app_scope":        tplBuildTemplateLogT.AppScope,
		"create_date":      utils.FormatDatetime(tplBuildTemplateLogT.CreateDate),
		"last_update_date": utils.FormatDatetime(tplBuildTemplateLogT.LastUpdateDate),
		"create_user_id":   tplBuildTemplateLogT.CreateUserId,
	}
	return respInfo
}
