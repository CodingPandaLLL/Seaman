package service

import (
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 *
 */
type SecurityService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.TplUserT, bool)
	GetByAdminId(adminId int64) (model.TplUserT, bool)
	//获取管理员总数
	GetAdminCount() (int64, error)
	SaveAvatarImg(adminId int64, fileName string) bool
	GetAdminList(offset, limit int) []*model.TplUserT
}

func NewSecurityService(db *xorm.Engine) SecurityService {
	return &securityService{
		engine: db,
	}
}

/**
 * 管理员的服务实现结构体
 */
type securityService struct {
	engine *xorm.Engine
}

/**
 * 查询管理员总数
 */
func (ac *securityService) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.TplUserT))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return count, nil
}

/**
 * 通过用户名和密码查询管理员
 */
func (ac *securityService) GetByAdminNameAndPassword(username, password string) (model.TplUserT, bool) {
	var user model.TplUserT

	ac.engine.Where(" account = ? and password = ? ", username, password).Get(&user)

	return user, user.Id != 0
}

/**
 * 查询管理员信息
 */
func (ac *securityService) GetByAdminId(adminId int64) (model.TplUserT, bool) {
	var user model.TplUserT

	ac.engine.Id(adminId).Get(&user)

	return user, user.Id != 0
}

/**
 * 保存头像信息
 */
func (ac *securityService) SaveAvatarImg(adminId int64, fileName string) bool {
	user := model.TplUserT{BackUp: fileName}
	_, err := ac.engine.Id(adminId).Cols(" avatar ").Update(&user)
	return err != nil
}

/**
 * 获取管理员列表
 * offset：获取管理员的便宜量
 * limit：请求管理员的条数
 */
func (ac *securityService) GetAdminList(offset, limit int) []*model.TplUserT {
	var adminList []*model.TplUserT

	err := ac.engine.Limit(limit, offset).Find(&adminList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return adminList
}
