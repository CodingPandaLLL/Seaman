package model

import (
	"Seaman/utils"
	"time"
)

type SmProjectT struct {
	Id               int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	ProjectName      string    `xorm:"not null comment('项目名称')  VARCHAR(255)"`
	ProjectCode      string    `xorm:"not null comment('项目编码')  VARCHAR(255)"`
	ProjectDesc      string    `xorm:"comment('项目描述')  VARCHAR(512)"`
	Status           string    `xorm:"comment('状态')  VARCHAR(2)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	LastUpdateUserId string    `xorm:"not null comment('更新用户ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间')"`
	LastUpdateDate   time.Time `xorm:"comment('更新时间')"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserName   string    `xorm:"extends"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (smProjectT *SmProjectT) ProjectTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":             smProjectT.Id,
		"projectDesc":    smProjectT.ProjectDesc,
		"projectName":    smProjectT.ProjectName,
		"projectCode":    smProjectT.ProjectCode,
		"status":         smProjectT.Status,
		"createUserId":   smProjectT.CreateUserId,
		"createUserName": smProjectT.CreateUserName,
		"createDate":     utils.FormatDatetime(smProjectT.CreateDate),
		"appName":        smProjectT.AppName,
		"tenantId":       smProjectT.TenantId,
		"appScope":       smProjectT.AppScope,
	}
	return respInfo
}
