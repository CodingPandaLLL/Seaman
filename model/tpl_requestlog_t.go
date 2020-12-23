package model

import (
	"time"
)

type TplRequestlogT struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	Path        string    `xorm:"comment('路径') VARCHAR(200)"`
	Method      string    `xorm:"comment('方法') VARCHAR(20)"`
	Body        string    `xorm:"comment('请求体') VARCHAR(2000)"`
	Status      int       `xorm:"comment('状态1') INT(11)"`
	RequestTime time.Time `xorm:"comment('请求时间') DATETIME"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplRequestlogT *TplRequestlogT) tplRequestlogTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           tplRequestlogT.Id,
		"path":         tplRequestlogT.Path,
		"method":       tplRequestlogT.Method,
		"body":         tplRequestlogT.Body,
		"status":       tplRequestlogT.Status,
		"request_time": tplRequestlogT.RequestTime,
	}
	return respInfo
}
