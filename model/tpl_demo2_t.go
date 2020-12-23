package model

import (
	"time"
)

type TplDemo2T struct {
	Id               int64     `xorm:"pk BIGINT(20)"`
	Id2              int64     `xorm:"not null pk BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     int64     `xorm:"comment('创建用户ID') BIGINT(20)"`
	LastUpdateUserId int64     `xorm:"comment('最后更新用户ID') BIGINT(20)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	Name             string    `xorm:"not null comment('姓名') VARCHAR(255)"`
	Age              int       `xorm:"not null comment('年龄') INT(20)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
	DespSt           string    `xorm:"comment('测试搜索引擎模糊查询') VARCHAR(512)"`
}
