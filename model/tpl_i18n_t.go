package model

import (
	"Seaman/utils"
	"time"
)

type TplI18nT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	KeyCode          string    `xorm:"comment('键代号') VARCHAR(255)"`
	LanguageCode     string    `xorm:"comment('语言代号') VARCHAR(255)"`
	ValueDesc        string    `xorm:"comment('值描述') VARCHAR(4000)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplI18nT *TplI18nT) tplI18nTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplI18nT.Id,
		"key_code":            tplI18nT.KeyCode,
		"language_code":       tplI18nT.LanguageCode,
		"value_desc":          tplI18nT.ValueDesc,
		"revision":            tplI18nT.Revision,
		"tenant_id":           tplI18nT.TenantId,
		"app_name":            tplI18nT.AppName,
		"app_scope":           tplI18nT.AppScope,
		"create_date":         utils.FormatDatetime(tplI18nT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplI18nT.LastUpdateDate),
		"create_user_id":      tplI18nT.CreateUserId,
		"last_update_user_id": tplI18nT.LastUpdateUserId,
	}
	return respInfo
}
