package service

import (
	"Seaman/config"
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"time"
)

/**
 * 角色模块功能服务接口
 */
type AppRoleService interface {

	//新增角色
	AddAppRole(model *model.TplAppRoleT) bool

	//删除角色
	DeleteAppRole(id int) bool

	//通过id查询角色
	GetAppRole(id int) (model.TplAppRoleT, error)

	//获取角色总数
	GetAppRoleTotalCount() (int64, error)

	//角色列表
	GetAppRoleList() []*model.TplAppRoleT

	//角色列表带分页
	GetAppRolePageList(appRole *model.TplAppRoleT,offset, limit int) []*model.TplAppRoleT

	//修改角色
	UpdateAppRole(model *model.TplAppRoleT) bool
}

/**
 * 实例化角色服务结构实体对象
 */
func NewAppRoleService(engine *xorm.Engine) AppRoleService {
	return &appRoleService{
		Engine: engine,
	}
}

/**
 * 角色服务实现结构体
 */
type appRoleService struct {
	Engine *xorm.Engine
}

/**
 *新增角色
 *appRole：角色信息
 */
func (ars *appRoleService) AddAppRole(appRole *model.TplAppRoleT) bool {

	//插入项目固定字段内容
	initConfig := config.InitConfig()
	appRole.AppName = initConfig.AppName
	appRole.AppScope = initConfig.AppScope
	appRole.TenantId = initConfig.TenantId
	appRole.CreateDate = time.Now()
	appRole.LastUpdateDate = time.Now()
	//todo:带入createUserId 和 UpdateUserId

	_, err := ars.Engine.Insert(appRole)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除角色
 */
func (ars *appRoleService) DeleteAppRole(appRoleId int) bool {

	appRole := model.TplAppRoleT{Id: int64(appRoleId)}
	_, err := ars.Engine.Where(" id = ? ", appRoleId).Delete(appRole)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询角色
 */
func (ars *appRoleService) GetAppRole(appRoleId int) (model.TplAppRoleT, error) {
	var appRole model.TplAppRoleT
	_, err := ars.Engine.Id(appRoleId).Get(&appRole)
	return appRole, err
}

/**
 * 请求总的角色数量
 * 返回值：总角色数量
 */
func (ars *appRoleService) GetAppRoleTotalCount() (int64, error) {

	count, err := ars.Engine.Where(" 1 = ? ", 1).Count(new(model.TplAppRoleT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//角色总数
	return count, nil
}

/**
 * 请求角色列表数据
 */
func (ars *appRoleService) GetAppRoleList() []*model.TplAppRoleT {

	var appRoleList []*model.TplAppRoleT
	err := ars.Engine.Where("1 = ?", 1).Find(&appRoleList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return appRoleList
}

/**
 * 请求角色列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (ars *appRoleService) GetAppRolePageList(appRole *model.TplAppRoleT,offset, limit int) []*model.TplAppRoleT {
	var appRoleList []*model.TplAppRoleT
	session := ars.Engine.Where("1 = ?", 1)
	err := session.Limit(limit, offset).Find(&appRoleList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return appRoleList
}

/**
 *修改角色
 *appRole：角色信息
 */
func (ars *appRoleService) UpdateAppRole(appRole *model.TplAppRoleT) bool {
	_, err := ars.Engine.Id(appRole.Id).Update(appRole)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
