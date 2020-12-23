package model

import (
	"time"
)

type TplTempT struct {
	Id               int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     int64     `xorm:"not null comment('创建人ID') BIGINT(20)"`
	LastUpdateUserId int64     `xorm:"not null comment('最后修改人ID') BIGINT(20)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	Name             string    `xorm:"VARCHAR(255)"`
	Test             time.Time `xorm:"DATE"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}
