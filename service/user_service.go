package service

import (
	"Seaman/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

/**
 * 用户模块功能服务接口
 */
type UserService interface {

	//新增用户
	AddUser(model *model.TplUserT) bool

	//删除用户
	DeleteUser(id int) bool

	//通过id查询用户
	GetUser(id int) (model.TplUserT, error)

	//获取用户总数
	GetUserTotalCount() (int64, error)

	//用户列表
	GetUserList() []*model.TplUserT

	//用户列表带分页
	GetUserPageList(offset, limit int) []*model.TplUserT

	//修改用户
	UpdateUser(model *model.TplUserT) bool
}

/**
 * 实例化用户服务结构实体对象
 */
func NewUserService(engine *xorm.Engine) UserService {
	return &userService{
		Engine: engine,
	}
}

/**
 * 用户服务实现结构体
 */
type userService struct {
	Engine *xorm.Engine
}

/**
 *新增用户
 *user：用户信息
 */
func (us *userService) AddUser(user *model.TplUserT) bool {
	//对密码进行加密
	pwd := user.Password
	m := md5.New()
	m.Write([]byte(pwd))
	pwd = hex.EncodeToString(m.Sum(nil))
	user.Password = pwd
	_, err := us.Engine.Insert(user)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除用户
 */
func (us *userService) DeleteUser(userId int) bool {

	user := model.TplUserT{Id: int64(userId)}
	_, err := us.Engine.Where(" id = ? ", userId).Delete(user)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 *通过id查询用户
 */
func (us *userService) GetUser(userId int) (model.TplUserT, error) {
	var user model.TplUserT
	_, err := us.Engine.Id(userId).Get(&user)
	return user, err
}

/**
 * 请求总的用户数量
 * 返回值：总用户数量
 */
func (uc *userService) GetUserTotalCount() (int64, error) {

	//查询del_flag 为0 的用户的总数量；del_flag:0 正常状态；del_flag:1 用户注销或者被删除
	count, err := uc.Engine.Where(" status = ? ", 1).Count(new(model.TplUserT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//用户总数
	return count, nil
}

/**
 * 请求用户列表数据
 */
func (uc *userService) GetUserList() []*model.TplUserT {

	var userList []*model.TplUserT
	err := uc.Engine.Where("status = ?", 1).Find(&userList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return userList
}

/**
 * 请求用户列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (uc *userService) GetUserPageList(offset, limit int) []*model.TplUserT {
	var userList []*model.TplUserT
	err := uc.Engine.Where("status = ?", 1).Limit(limit, offset).Find(&userList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return userList
}

/**
 *修改用户
 *user：用户信息
 */
func (us *userService) UpdateUser(user *model.TplUserT) bool {
	_, err := us.Engine.Id(user.Id).Update(user)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
