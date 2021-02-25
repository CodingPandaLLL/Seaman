package service

import (
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

/**
 * 项目模块功能服务接口
 */
type ProjectService interface {

	//新增项目
	AddProject(model *model.SmProjectT) bool

	//删除项目
	DeleteProject(id int) bool

	//通过id查询项目
	GetProject(id int) (model.SmProjectT, error)

	//获取项目总数
	GetProjectTotalCount(project *model.SmProjectT) (int64, error)

	//项目列表
	GetProjectList(project *model.SmProjectT) []*model.SmProjectT

	//项目列表带分页
	GetProjectPageList(project *model.SmProjectT, offset, limit int) []*model.SmProjectT

	//修改项目
	UpdateProject(model *model.SmProjectT) bool
}

/**
 * 实例化项目服务结构实体对象
 */
func NewProjectService(engine *xorm.Engine) ProjectService {
	return &projectService{
		Engine: engine,
	}
}

/**
 * 项目服务实现结构体
 */
type projectService struct {
	Engine *xorm.Engine
}

/**
 *新增项目
 *project：项目信息
 */
func (ps *projectService) AddProject(project *model.SmProjectT) bool {

	_, err := ps.Engine.Insert(project)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除项目
 */
func (ps *projectService) DeleteProject(projectId int) bool {

	project := model.SmProjectT{Id: int64(projectId)}
	_, err := ps.Engine.Where(" id = ? ", projectId).Delete(project)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询项目
 */
func (ps *projectService) GetProject(projectId int) (model.SmProjectT, error) {
	var project model.SmProjectT
	_, err := ps.Engine.Id(projectId).Get(&project)
	return project, err
}

/**
 * 请求总的项目数量
 * 返回值：总项目数量
 */
func (ps *projectService) GetProjectTotalCount(project *model.SmProjectT) (int64, error) {
	session := ps.Engine.Where("1 = ?", 1)
	if len(project.Status)!=0 {
		session.And("status = ?",project.Status)
	}
	if len(project.ProjectCode)!=0 {
		projectCode := "%"+project.ProjectCode+"%"
		session.And("project_code like ?",projectCode)
	}
	if len(project.ProjectName)!=0 {
		projectName := "%"+project.ProjectName+"%"
		session.And("project_name like ?",projectName)
	}
	count, err := session.Count(new(model.SmProjectT))
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
func (ps *projectService) GetProjectList(project *model.SmProjectT) []*model.SmProjectT {
	var projectList []*model.SmProjectT
	//err := ps.Engine.Where("1 = ?", 1).Find(&projectList)
	session := ps.Engine.Where(" 1 = ? ", 1)
	if project.Id != 0 {
		session.And("group_id = ?", project.Id)
	}
	err := session.Find(&projectList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return projectList
}

/**
 * 请求项目列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (ps *projectService) GetProjectPageList(project *model.SmProjectT, offset, limit int) []*model.SmProjectT {
	var projectList []*model.SmProjectT
	session := ps.Engine.Where("1 = ?", 1)
	if len(project.Status)!=0 {
		session.And("status = ?",project.Status)
	}
	if len(project.ProjectCode)!=0 {
		projectCode := "%"+project.ProjectCode+"%"
		session.And("project_code like ?",projectCode)
	}
	if len(project.ProjectName)!=0 {
		projectName := "%"+project.ProjectName+"%"
		session.And("project_name like ?",projectName)
	}
	err := session.Limit(limit, offset).Find(&projectList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	//给CreateUserName赋值
	for i := 0; i < len(projectList); i++ {
		// todo 根据用户id查询用户信息，后期替换为redis
		var user model.TplUserT
		_, userErr := ps.Engine.Id(projectList[i].CreateUserId).Get(&user)
		if userErr != nil {
			iris.New().Logger().Error(userErr.Error())
			panic(userErr.Error())
			return nil
		}
		projectList[i].CreateUserName = user.Lastname
	}
	return projectList
}

/**
 *修改项目
 *project：项目信息
 */
func (ps *projectService) UpdateProject(project *model.SmProjectT) bool {
	_, err := ps.Engine.Id(project.Id).Update(project)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
