package service

import (
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

/**
 * 群组用户管理模块功能服务接口
 */
type UserGroupService interface {

	//新增群组用户管理
	AddUserGroup(model *model.TplUserGroupT) bool

	//删除群组用户管理
	DeleteUserGroup(id int) bool

	//通过id查询群组用户管理
	GetUserGroup(id int) (model.TplUserGroupT, error)

	//获取群组用户管理总数
	GetUserGroupTotalCount() (int64, error)

	//群组用户管理列表
	GetUserGroupList(userGroup *model.TplUserGroupT) []*model.TplUserGroupT

	//群组用户管理列表带分页
	GetUserGroupPageList(userGroup *model.TplUserGroupT,offset, limit int) []*model.TplUserGroupT

	//修改群组用户管理
	UpdateUserGroup(model *model.TplUserGroupT) bool
}

/**
 * 实例化群组用户管理服务结构实体对象
 */
func NewUserGroupService(engine *xorm.Engine) UserGroupService {
	return &userGroupService{
		Engine: engine,
	}
}

/**
 * 群组用户管理服务实现结构体
 */
type userGroupService struct {
	Engine *xorm.Engine
}

/**
 *新增群组用户管理
 *userGroup：群组用户管理信息
 */
func (ugr *userGroupService) AddUserGroup(userGroup *model.TplUserGroupT) bool {

	_, err := ugr.Engine.Insert(userGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除群组用户管理
 */
func (ugr *userGroupService) DeleteUserGroup(userGroupId int) bool {

	userGroup := model.TplUserGroupT{GroupId: int64(userGroupId)}
	_, err := ugr.Engine.Where(" id = ? ", userGroupId).Delete(userGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询群组用户管理
 */
func (ugr *userGroupService) GetUserGroup(userGroupId int) (model.TplUserGroupT, error) {
	var userGroup model.TplUserGroupT
	_, err := ugr.Engine.Id(userGroupId).Get(&userGroup)
	return userGroup, err
}

/**
 * 请求总的群组用户管理数量
 * 返回值：总群组用户管理数量
 */
func (ugr *userGroupService) GetUserGroupTotalCount() (int64, error) {

	count, err := ugr.Engine.Where(" 1 = ? ", 1).Count(new(model.TplUserGroupT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//群组用户管理总数
	return count, nil
}

/**
 * 请求群组用户管理列表数据
 */
func (ugr *userGroupService) GetUserGroupList(userGroup *model.TplUserGroupT) []*model.TplUserGroupT {
	var userGroupList []*model.TplUserGroupT
	//err := ugr.Engine.Where("1 = ?", 1).Find(&userGroupList)
	session := ugr.Engine.Where(" 1 = ? ", 1)
	if  userGroup.GroupId!=0 {
		session.And("group_id = ?",userGroup.GroupId)
	}
	if  userGroup.UserId!=0 {
		session.And("user_id = ?",userGroup.UserId)
	}
	err := session.Find(&userGroupList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return userGroupList
}

/**
 * 请求群组用户管理列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (ugr *userGroupService) GetUserGroupPageList(userGroup *model.TplUserGroupT,offset, limit int) []*model.TplUserGroupT {
	var userGroupList []*model.TplUserGroupT
	session := ugr.Engine.Where("1 = ?", 1)
	err := session.Limit(limit, offset).Find(&userGroupList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return userGroupList
}

/**
 *修改群组用户管理
 *userGroup：群组用户管理信息
 */
func (ugr *userGroupService) UpdateUserGroup(userGroup *model.TplUserGroupT) bool {
	_, err := ugr.Engine.Id(userGroup.GroupId).Update(userGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
