package service

import (
	"Seaman/config"
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
	"time"
)

/**
 * 群组模块功能服务接口
 */
type AppGroupService interface {

	//新增群组
	AddAppGroup(model *model.TplAppGroupT) bool

	//删除群组
	DeleteAppGroup(id int) bool

	//通过id查询群组
	GetAppGroup(id int) (model.TplAppGroupT, error)

	//获取群组总数
	GetAppGroupTotalCount(appGroup *model.TplAppGroupT) (int64, error)

	//查询群组成员列表
	GetGroupMembers(id int) []*model.TplUserT

	//群组列表
	GetAppGroupList() []*model.TplAppGroupT

	//群组列表带分页
	GetAppGroupPageList(appGroup *model.TplAppGroupT, offset, limit int) []*model.TplAppGroupT

	//修改群组
	UpdateAppGroup(model *model.TplAppGroupT) bool

	//更新群组信息
	UpdateGroupMembers(id int64, ids string) bool
}

/**
 * 实例化群组服务结构实体对象
 */
func NewAppGroupService(engine *xorm.Engine) AppGroupService {
	return &appGroupService{
		Engine: engine,
	}
}

/**
 * 群组服务实现结构体
 */
type appGroupService struct {
	Engine *xorm.Engine
}

/**
 *新增群组
 *appGroup：群组信息
 */
func (ags *appGroupService) AddAppGroup(appGroup *model.TplAppGroupT) bool {

	//插入项目固定字段内容
	initConfig := config.InitConfig()
	appGroup.AppName = initConfig.AppName
	appGroup.AppScope = initConfig.AppScope
	appGroup.TenantId = initConfig.TenantId
	appGroup.CreateDate = time.Now()
	appGroup.LastUpdateDate = time.Now()
	//todo:带入createUserId 和 UpdateUserId

	_, err := ags.Engine.Insert(appGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除群组
 */
func (ags *appGroupService) DeleteAppGroup(appGroupId int) bool {

	//删除已存在的关联关系
	userGroup := model.TplUserGroupT{GroupId: int64(appGroupId)}
	_, userGroupErr := ags.Engine.Delete(userGroup)
	if userGroupErr != nil {
		iris.New().Logger().Info(userGroupErr.Error())
	}

	//删除群组
	appGroup := model.TplAppGroupT{Id: int64(appGroupId)}
	_, err := ags.Engine.Where(" id = ? ", appGroupId).Delete(appGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询群组
 */
func (ags *appGroupService) GetAppGroup(appGroupId int) (model.TplAppGroupT, error) {
	var appGroup model.TplAppGroupT
	_, err := ags.Engine.Id(appGroupId).Get(&appGroup)
	return appGroup, err
}

/**
 * 请求总的群组数量
 * 返回值：总群组数量
 */
func (ags *appGroupService) GetAppGroupTotalCount(appGroup *model.TplAppGroupT) (int64, error) {
	session := ags.Engine.Where(" 1 = ? ", 1)
	if appGroup.Name != "" && len(appGroup.Name) > 0 {
		name := "%" + appGroup.Name + "%"
		session.And("name like ?", name)
	}
	if appGroup.Code != "" && len(appGroup.Code) > 0 {
		code := "%" + appGroup.Code + "%"
		session.And("code like ?", code)
	}
	if appGroup.Status != 0 {
		session.And("status = ?", appGroup.Status)
	}
	if appGroup.Type != "" && len(appGroup.Type) > 0 {
		sType := "%" + appGroup.Type + "%"
		session.And("type like ?", sType)
	}
	count, err := session.Count(new(model.TplAppGroupT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//群组总数
	return count, nil
}

/**
 * 请求群组列表数据
 */
func (ags *appGroupService) GetAppGroupList() []*model.TplAppGroupT {

	var appGroupList []*model.TplAppGroupT
	err := ags.Engine.Where("1 = ?", 1).Find(&appGroupList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return appGroupList
}

/**
 * 请求群组成员数据
 */
func (ags *appGroupService) GetGroupMembers(groupId int) []*model.TplUserT {

	//查询群组关联
	var userGroupList []*model.TplUserGroupT
	err := ags.Engine.Where("group_id = ?", groupId).Find(&userGroupList)

	//查询用户
	var userList []*model.TplUserT
	userErr := ags.Engine.Where("status = ?", 1).Find(&userList)
	if userErr != nil {
		iris.New().Logger().Error(userErr.Error())
		panic(userErr.Error())
		return nil
	}

	//数据处理
	for i := 0; i < len(userList); i++ {
		userId := userList[i].Id
		for j := 0; j < len(userGroupList); j++ {
			if userId == userGroupList[j].UserId {
				userList[i].BackUp = "chosen"
				break
			}
		}
	}

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return userList
}

/**
 * 请求群组列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (ags *appGroupService) GetAppGroupPageList(appGroup *model.TplAppGroupT, offset, limit int) []*model.TplAppGroupT {
	var appGroupList []*model.TplAppGroupT
	session := ags.Engine.Where("1 = ?", 1)
	if appGroup.Name != "" && len(appGroup.Name) > 0 {
		name := "%" + appGroup.Name + "%"
		session.And("name like ?", name)
	}
	if appGroup.Code != "" && len(appGroup.Code) > 0 {
		code := "%" + appGroup.Code + "%"
		session.And("code like ?", code)
	}
	if appGroup.Status != 0 {
		session.And("status = ?", appGroup.Status)
	}
	if appGroup.Type != "" && len(appGroup.Type) > 0 {
		sType := "%" + appGroup.Type + "%"
		session.And("type like ?", sType)
	}
	err := session.Limit(limit, limit*offset).Find(&appGroupList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return appGroupList
}

/**
 *修改群组
 *appGroup：群组信息
 */
func (ags *appGroupService) UpdateAppGroup(appGroup *model.TplAppGroupT) bool {
	_, err := ags.Engine.Id(appGroup.Id).Update(appGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 *修改群组
 *appGroup：群组信息
 */
func (ags *appGroupService) UpdateGroupMembers(groupId int64, userIds string) bool {
	//删除已存在的关联关系
	userGroup := model.TplUserGroupT{GroupId: groupId}
	_, err := ags.Engine.Delete(userGroup)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	//新增关联关系
	sep := ","
	arr := strings.Split(userIds, sep)
	for i:=0;i< len(arr);i++ {
		userId, strErr := strconv.ParseInt(arr[i], 10, 64)
		if strErr != nil {
			iris.New().Logger().Info(strErr.Error())
		}
		newUserGroup := model.TplUserGroupT{GroupId: groupId,UserId: userId}
		_, userGroupErr := ags.Engine.Insert(newUserGroup)
		if userGroupErr != nil {
			iris.New().Logger().Info(userGroupErr.Error())
		}
	}
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
