package model

type TplDataColRuleT struct {
	Id         int    `xorm:"not null pk INT(11)"`
	Code       string `xorm:"not null comment('编码') VARCHAR(20)"`
	DataFileds string `xorm:"comment('有效字段集合') VARCHAR(16000)"`
	DataRuleId int    `xorm:"not null comment('数据权限id') INT(11)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplDataColRuleT *TplDataColRuleT) tplDataColRuleTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           tplDataColRuleT.Id,
		"code":         tplDataColRuleT.Code,
		"data_fileds":  tplDataColRuleT.DataFileds,
		"data_rule_id": tplDataColRuleT.DataRuleId,
	}
	return respInfo
}
