package model

import (
	"Seaman/utils"
	"time"
)

type TplLookupT struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	Code             string    `xorm:"not null comment('编号') unique(LOOKUP_CODE_INDEX) VARCHAR(100)"`
	Desp             string    `xorm:"not null comment('描述') VARCHAR(100)"`
	Type             string    `xorm:"comment('类型') VARCHAR(50)"`
	ParentId         int64     `xorm:"not null comment('父节点ID') index BIGINT(20)"`
	Revision         int64     `xorm:"not null comment('版本号') BIGINT(20)"`
	CreateUserId     string    `xorm:"not null comment('创建用户ID') VARCHAR(36)"`
	CreateDate       time.Time `xorm:"comment('创建时间') DATETIME"`
	LastUpdateUserId string    `xorm:"not null comment('最后更新用户ID') VARCHAR(36)"`
	LastUpdateDate   time.Time `xorm:"comment('最后修改时间') DATETIME"`
	GroupCode        string    `xorm:"not null comment('群组编号') unique(LOOKUP_CODE_INDEX) VARCHAR(100)"`
	AppName          string    `xorm:"not null comment('应用名') VARCHAR(32)"`
	TenantId         string    `xorm:"comment('多租户ID') VARCHAR(32)"`
	AppScope         string    `xorm:"comment('系统群名') VARCHAR(32)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplLookupT *TplLookupT) tplLookupTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplLookupT.Id,
		"code":                tplLookupT.Code,
		"desp":                tplLookupT.Desp,
		"type":                tplLookupT.Type,
		"parent_id":           tplLookupT.ParentId,
		"revision":            tplLookupT.Revision,
		"tenant_id":           tplLookupT.TenantId,
		"app_name":            tplLookupT.AppName,
		"app_scope":           tplLookupT.AppScope,
		"create_date":         utils.FormatDatetime(tplLookupT.CreateDate),
		"last_update_date":    utils.FormatDatetime(tplLookupT.LastUpdateDate),
		"create_user_id":      tplLookupT.CreateUserId,
		"last_update_user_id": tplLookupT.LastUpdateUserId,
	}
	return respInfo
}
