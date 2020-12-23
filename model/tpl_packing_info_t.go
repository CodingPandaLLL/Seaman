package model

import (
	"time"
)

type TplPackingInfoT struct {
	Id               int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Name             string    `xorm:"not null comment('模板名称') VARCHAR(100)"`
	Version          string    `xorm:"not null comment('模板版本号') VARCHAR(80)"`
	Desc             string    `xorm:"comment('描述') VARCHAR(250)"`
	UploadUserId     string    `xorm:"not null comment('模板上传用户') VARCHAR(40)"`
	UploadTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('模板上传时间') TIMESTAMP"`
	LastUpdateUserId string    `xorm:"not null comment('更新人ID') VARCHAR(40)"`
	LastUpdateTime   time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	DownloadUrl      string    `xorm:"comment('模板文件的下载地址') VARCHAR(250)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tplPackingInfoT *TplPackingInfoT) tplPackingInfoTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                  tplPackingInfoT.Id,
		"name":                tplPackingInfoT.Name,
		"desc":                tplPackingInfoT.Desc,
		"upload_user_id":      tplPackingInfoT.UploadUserId,
		"upload_time":         tplPackingInfoT.UploadTime,
		"download_url":        tplPackingInfoT.DownloadUrl,
		"last_update_user_id": tplPackingInfoT.LastUpdateUserId,
	}
	return respInfo
}
