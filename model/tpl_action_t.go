package model

import (
	"Seaman/utils"
	"time"
)

type TplActionT struct {
	Id           int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	OperaCode    int       `xorm:"index(FILTER) comment('操作码') INT(11)"`
	OperaDesc    string    `xorm:"index(FILTER) comment('操作描述')  VARCHAR(255)"`
	ResCode      int       `xorm:"index(FILTER) comment('资源码') INT(11)"`
	ResDesc      string    `xorm:"index(FILTER) comment('资源描述')  VARCHAR(255)"`
	ParamDesc    string    `xorm: comment('参数') "VARCHAR(4000)"`
	CreateUserId string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	CreateDate   time.Time `xorm:"comment('创建时间') index DATETIME"`
	IpAddress    string    `xorm:"index  comment('ip地址')  VARCHAR(80)"`
	MethodName   string    `xorm:"index(FILTER)  comment('方法名')  VARCHAR(255)"`
	AppName      string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId     string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope     string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplActionT *TplActionT) tplActionTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":             tplActionT.Id,
		"opera_code":     tplActionT.OperaCode,
		"opera_desc":     tplActionT.OperaDesc,
		"res_code":       tplActionT.ResCode,
		"res_desc":       tplActionT.ResDesc,
		"param_desc":     tplActionT.ParamDesc,
		"create_user_id": tplActionT.CreateUserId,
		"create_date":    utils.FormatDatetime(tplActionT.CreateDate),
		"ip_address":     tplActionT.IpAddress,
		"method_name":    tplActionT.MethodName,
		"app_name":       tplActionT.AppName,
		"tenant_id":      tplActionT.TenantId,
		"app_scope":      tplActionT.AppScope,
	}
	return respInfo
}
