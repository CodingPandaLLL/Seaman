package model

type SmProjectFileT struct {
	FileId    int64 `xorm:"not null  comment('文件ID')  BIGINT(20)"`
	ProjectId int64 `xorm:"not null  comment('项目ID')  BIGINT(20)"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (tmProjectFileT *SmProjectFileT) ProjectFileTToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"projectId": tmProjectFileT.ProjectId,
		"fileId":    tmProjectFileT.FileId,
	}
	return respInfo
}
