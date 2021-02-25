package service

import (
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

/**
 * 项目模块功能服务接口
 */
type ProjectFileService interface {

	//新增项目
	AddProjectFile(model *model.SmProjectFileT) bool

	//删除项目
	DeleteProjectFile(id int) bool

	//通过id查询项目
	GetProjectFile(fileId int) (model.TplFileT, error)

	//获取项目总数
	GetProjectFileTotalCount(projectFile *model.SmProjectFileT) (int64, error)

	//项目列表
	GetProjectFileList(projectFile *model.SmProjectFileT) []*model.TplFileT

	//项目列表带分页
	GetProjectFilePageList(projectFile *model.SmProjectFileT, offset, limit int) []*model.SmProjectFileT

	//修改项目
	UpdateProjectFile(model *model.SmProjectFileT) bool
}

/**
 * 实例化项目服务结构实体对象
 */
func NewProjectFileService(engine *xorm.Engine) ProjectFileService {
	return &projectFileService{
		Engine: engine,
	}
}

/**
 * 项目服务实现结构体
 */
type projectFileService struct {
	Engine *xorm.Engine
}

/**
 *新增项目
 *projectFile：项目信息
 */
func (pfs *projectFileService) AddProjectFile(projectFile *model.SmProjectFileT) bool {

	_, err := pfs.Engine.Insert(projectFile)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除项目
 */
func (pfs *projectFileService) DeleteProjectFile(projectFileId int) bool {

	projectFile := model.SmProjectFileT{ProjectId: int64(projectFileId)}
	_, err := pfs.Engine.Where(" id = ? ", projectFileId).Delete(projectFile)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询项目
 */
func (pfs *projectFileService) GetProjectFile(fileId int) (model.TplFileT, error) {
	var projectFile model.TplFileT
	_, err := pfs.Engine.Id(fileId).Get(&projectFile)
	return projectFile, err
}

/**
 * 请求总的项目数量
 * 返回值：总项目数量
 */
func (pfs *projectFileService) GetProjectFileTotalCount(projectFile *model.SmProjectFileT) (int64, error) {
	session := pfs.Engine.Where("1 = ?", 1)
	count, err := session.Count(new(model.SmProjectFileT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//项目总数
	return count, nil
}

/**
 * 请求项目列表数据
 */
func (pfs *projectFileService) GetProjectFileList(projectFile *model.SmProjectFileT) []*model.TplFileT {
	var fileList []*model.TplFileT
	//多表关联查询
	session := pfs.Engine.SQL("SELECT tft.* FROM seaman.tpl_file_t tft left join seaman.sm_project_file_t spft on spft.file_id =tft.id where spft.project_id= ?",projectFile.ProjectId)
	err := session.Find(&fileList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return fileList
}

/**
 * 请求项目列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (pfs *projectFileService) GetProjectFilePageList(projectFile *model.SmProjectFileT, offset, limit int) []*model.SmProjectFileT {
	var projectFileList []*model.SmProjectFileT
	session := pfs.Engine.Where("1 = ?", 1)
	err := session.Limit(limit, offset).Find(&projectFileList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return projectFileList
}

/**
 *修改项目
 *projectFile：项目信息
 */
func (pfs *projectFileService) UpdateProjectFile(projectFile *model.SmProjectFileT) bool {
	_, err := pfs.Engine.Id(projectFile.FileId).Update(projectFile)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
