package model

type TplDataRowRuleT struct {
	Id           int    `xorm:"not null pk INT(11)"`
	Code         string `xorm:"not null comment('编码') VARCHAR(20)"`
	DefaultField string `xorm:"not null comment('默认作用于表字段名') VARCHAR(20)"`
	DataType     string `xorm:"not null comment('数据类型') VARCHAR(20)"`
	SqlOperator  string `xorm:"not null comment('操作符') VARCHAR(10)"`
	DataValues   string `xorm:"comment('配置值') VARCHAR(16000)"`
	DataRuleId   int    `xorm:"not null comment('数据权限id') INT(11)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplDataRowRuleT *TplDataRowRuleT) tplDataRowRuleTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":            tplDataRowRuleT.Id,
		"code":          tplDataRowRuleT.Code,
		"default_field": tplDataRowRuleT.DefaultField,
		"data_type":     tplDataRowRuleT.DataType,
		"sql_operator":  tplDataRowRuleT.SqlOperator,
		"data_values":   tplDataRowRuleT.DataValues,
		"data_rule_id":  tplDataRowRuleT.DataRuleId,
	}
	return respInfo
}
