package service

import (
	"Seaman/config"
	"Seaman/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"time"
)

/**
 * 组织机构模块功能服务接口
 */
type OrgService interface {

	//新增组织机构
	AddOrg(model *model.TplOrgT) bool

	//删除组织机构
	DeleteOrg(id int) bool

	//删除组织机构的子节点
	DeleteOrgChildren(pId int) bool

	//通过id查询组织机构
	GetOrg(id int) (model.TplOrgT, error)

	//获取组织机构总数
	GetOrgTotalCount() (int64, error)

	//组织机构列表
	GetOrgList() []*model.TplOrgT

	//组织机构列表带分页
	GetOrgPageList(org *model.TplOrgT,offset, limit int) []*model.TplOrgT

	//修改组织机构
	UpdateOrg(model *model.TplOrgT) bool
}

/**
 * 实例化组织机构服务结构实体对象
 */
func NewOrgService(engine *xorm.Engine) OrgService {
	return &orgService{
		Engine: engine,
	}
}

/**
 * 组织机构服务实现结构体
 */
type orgService struct {
	Engine *xorm.Engine
}

/**
 *新增组织机构
 *org：组织机构信息
 */
func (os *orgService) AddOrg(org *model.TplOrgT) bool {

	//插入项目固定字段内容
	initConfig := config.InitConfig()
	org.AppName = initConfig.AppName
	org.AppScope = initConfig.AppScope
	org.TenantId = initConfig.TenantId
	org.CreateDate = time.Now()
	org.LastUpdateDate = time.Now()
	org.Revision = 1
	//todo:带入createUserId 和 UpdateUserId

	_, err := os.Engine.Insert(org)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除组织机构
 */
func (os *orgService) DeleteOrg(orgId int) bool {

	org := model.TplOrgT{Id: int64(orgId)}
	_, err := os.Engine.Where(" id = ? ", orgId).Delete(org)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 * 删除组织机构子节点
 */
func (os *orgService) DeleteOrgChildren(pId int) bool {
	org := model.TplOrgT{}
	_, err := os.Engine.Where(" P_ID = ? ", pId).Delete(org)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 *通过id查询组织机构
 */
func (os *orgService) GetOrg(orgId int) (model.TplOrgT, error) {
	var org model.TplOrgT
	_, err := os.Engine.Id(orgId).Get(&org)
	return org, err
}

/**
 * 请求总的组织机构数量
 * 返回值：总组织机构数量
 */
func (os *orgService) GetOrgTotalCount() (int64, error) {

	count, err := os.Engine.Where(" 1 = ? ", 1).Count(new(model.TplOrgT))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	//组织机构总数
	return count, nil
}

/**
 * 请求组织机构列表数据
 */
func (os *orgService) GetOrgList() []*model.TplOrgT {

	var orgList []*model.TplOrgT
	err := os.Engine.Where("1 = ?", 1).Find(&orgList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return orgList
}

/**
 * 请求组织机构列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (os *orgService) GetOrgPageList(org *model.TplOrgT,offset, limit int) []*model.TplOrgT {
	var orgList []*model.TplOrgT
	session := os.Engine.Where("1 = ?", 1)
	err := session.Limit(limit, offset).Find(&orgList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return orgList
}

/**
 *修改组织机构
 *org：组织机构信息
 */
func (os *orgService) UpdateOrg(org *model.TplOrgT) bool {
	_, err := os.Engine.Id(org.Id).Update(org)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}
