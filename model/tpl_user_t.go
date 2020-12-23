package model

import (
	"Seaman/utils"
	"time"
)

type TplUserT struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Account        string    `xorm:"not null comment('帐号') unique unique VARCHAR(80)"`
	Email          string    `xorm:"comment('邮箱') VARCHAR(80)"`
	Firstname      string    `xorm:"comment('姓') VARCHAR(80)"`
	Lastname       string    `xorm:"not null comment('名') VARCHAR(80)"`
	Status         string    `xorm:"comment('状态') VARCHAR(3)"`
	Addr1          string    `xorm:"comment('地址1') VARCHAR(80)"`
	Addr2          string    `xorm:"comment('地址2') VARCHAR(40)"`
	City           string    `xorm:"comment('所在城市') VARCHAR(80)"`
	State          string    `xorm:"comment('所在州省') VARCHAR(80)"`
	Zip            string    `xorm:"comment('所在邮编') VARCHAR(20)"`
	Country        string    `xorm:"comment('所在国家') VARCHAR(20)"`
	Phone          string    `xorm:"comment('联系方式') VARCHAR(80)"`
	LanguageCode   string    `xorm:"not null default 'cn' comment('语言') VARCHAR(40)"`
	Password       string    `xorm:"not null comment('密码') VARCHAR(128)"`
	Organization   string    `xorm:"not null comment('所在部门') VARCHAR(100)"`
	Defaultin      string    `xorm:"not null default '0' comment('系统内置标识（0否1是）') CHAR(1)"`
	BackUp         string    `xorm:"comment('备用字段') VARCHAR(255)"`
	BackUp1        string    `xorm:"comment('备用字段1') VARCHAR(255)"`
	BackUp2        string    `xorm:"comment('备用字段2') VARCHAR(255)"`
	TenantId       string    `xorm:"comment('租户') VARCHAR(32)"`
	AppName        string    `xorm:"not null comment('系统名称') VARCHAR(32)"`
	AppScope       string    `xorm:"comment('系统群') VARCHAR(32)"`
	CreateDate     time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateDate time.Time `xorm:"comment('最后修改时间') DATETIME"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *TplUserT) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":               user.Id,
		"account":          user.Account,
		"email":            user.Email,
		"firstname":        user.Firstname,
		"lastname":         user.Lastname,
		"status":           user.Status,
		"addr1":            user.Addr1,
		"addr2":            user.Addr2,
		"city":             user.City,
		"state":            user.State,
		"zip":              user.Zip,
		"country":          user.Country,
		"language_code":    user.LanguageCode,
		"password":         user.Password,
		"organization":     user.Organization,
		"tenant_id":        user.TenantId,
		"app_name":         user.AppName,
		"app_scope":        user.AppScope,
		"create_date":      utils.FormatDatetime(user.CreateDate),
		"last_update_date": utils.FormatDatetime(user.LastUpdateDate),
	}
	return respInfo
}
