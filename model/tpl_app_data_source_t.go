package model

import (
	"Seaman/utils"
	"time"
)

type TplAppDataSourceT struct {
	Id               int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Type             string    `xorm:"comment('类型(0行1列)') VARCHAR(8)"`
	Code             string    `xorm:"comment('编码名称') VARCHAR(128)"`
	DataType         string    `xorm:"comment('字段类型') VARCHAR(32)"`
	SqlOperator      string    `xorm:"comment('操作') VARCHAR(16)"`
	DefaultField     string    `xorm:"comment('字段名') VARCHAR(255)"`
	Source           string    `xorm:"comment('源地址') VARCHAR(255)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
	CreateDate       time.Time `xorm:"not null comment('创建时间') DATETIME"`
	LastUpdateDate   time.Time `xorm:"not null comment('最后更新时间') DATETIME"`
	CreateUserId     string    `xorm:"not null comment('创建人') VARCHAR(32)"`
	LastUpdateUserId string    `xorm:"not null comment('更新人') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplAppDataSourceT *TplAppDataSourceT) tplAppDataSourceTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplAppDataSourceT.Id,
		"type":                tplAppDataSourceT.Type,
		"code":                tplAppDataSourceT.Code,
		"data_type":           tplAppDataSourceT.DataType,
		"sql_operator":        tplAppDataSourceT.SqlOperator,
		"default_field":       tplAppDataSourceT.DefaultField,
		"source":              tplAppDataSourceT.Source,
		"tenant_id":           tplAppDataSourceT.TenantId,
		"app_name":            tplAppDataSourceT.AppName,
		"app_scope":           tplAppDataSourceT.AppScope,
		"create_date":         utils.FormatDatetime(tplAppDataSourceT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplAppDataSourceT.LastUpdateDate),
		"create_user_id":      tplAppDataSourceT.CreateUserId,
		"last_update_user_id": tplAppDataSourceT.LastUpdateUserId,
	}
	return respInfo
}
