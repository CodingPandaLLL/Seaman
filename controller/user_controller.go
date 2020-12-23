package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
	"time"
)

//每一页最大的内容
const MaxLimit = 50

/**
 * 用户控制器结构体：用来实现处理用户模块的接口的请求，并返回给客户端
 */
type UserController struct {
	//上下文对象
	Ctx iris.Context
	//user service
	UserService service.UserService
	//session对象
	Session *sessions.Session
}

func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {

	//添加用户
	a.Handle("POST", "/", "PostAddUser")

	//删除用户
	a.Handle("DELETE", "/{id}", "DeleteUser")

	//查询用户
	a.Handle("GET", "/{id}", "GetUser")

	//查询用户数量
	a.Handle("GET", "/count", "GetCount")

	//所有用户列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询用户分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改用户
	a.Handle("PUT", "/", "UpdateUser")
}

/**
 * url: /user/adduser
 * type：post
 * descs：添加用户
 */
func (uc *UserController) PostAddUser() mvc.Result {

	var user model.TplUserT
	err := uc.Ctx.ReadJSON(&user)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	newUser := &model.TplUserT{
		Account:        user.Account,
		Email:          user.Email,
		Firstname:      user.Firstname,
		Lastname:       user.Lastname,
		Status:         user.Status,
		Addr1:          user.Addr1,
		Addr2:          user.Addr2,
		City:           user.City,
		State:          user.State,
		Zip:            user.Zip,
		Country:        user.Country,
		LanguageCode:   user.LanguageCode,
		Password:       user.Password,
		Organization:   user.Organization,
		TenantId:       user.TenantId,
		AppName:        user.AppName,
		AppScope:       user.AppScope,
		CreateDate:     time.Now(),
		LastUpdateDate: time.Now(),
	}
	isSuccess := uc.UserService.AddUser(newUser)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SUCCESS_ADD),
		},
	}
}

/**
 * 删除用户
 */
func (uc *UserController) DeleteUser() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	userId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := uc.UserService.DeleteUser(userId)
	if !delete {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELETE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELETE),
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_OK,
				"type":    utils.RESPMSG_SUCCESS_DELETE,
				"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELETE),
			},
		}
	}
}

/**
 * 获取用户
 * 请求类型：Get
 */
func (uc *UserController) GetUser() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	userId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回用户
	return mvc.Response{
		Object: user,
	}
}

/**
 * 获取用户数量
 * 请求类型：Get
 */
func (uc *UserController) GetCount() mvc.Result {

	//用户总数
	total, err := uc.UserService.GetUserTotalCount()

	//请求出现错误
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//正常情况的返回值
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  total,
		},
	}
}

/**
 * 获取用户列表
 * 请求类型：Get
 */
func (uc *UserController) GetList() mvc.Result {

	userList := uc.UserService.GetUserList()
	if len(userList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, user := range userList {
		respList = append(respList, user.UserToRespDesc())
	}

	//返回用户列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取用户带参数分页查询
 * 请求类型：Get
 */
func (uc *UserController) GetPageList() mvc.Result {

	offsetStr := uc.Ctx.FormValue("offset")
	limitStr := uc.Ctx.FormValue("limit")
	var offset int
	var limit int

	//判断offset和limit两个变量任意一个都不能为""
	if offsetStr == "" || limitStr == "" {

		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	offset, err := strconv.Atoi(offsetStr)
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	//做页数的限制检查
	if offset <= 0 {
		offset = 0
	}

	//做最大的限制
	if limit > MaxLimit {
		limit = MaxLimit
	}

	userList := uc.UserService.GetUserPageList(offset, limit)

	if len(userList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, user := range userList {
		respList = append(respList, user.UserToRespDesc())
	}

	//返回用户列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * type：put
 * descs：修改用户
 */
func (uc *UserController) UpdateUser() mvc.Result {

	var user model.TplUserT
	err := uc.Ctx.ReadJSON(&user)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}
	newUser := &model.TplUserT{
		Id:             user.Id,
		Account:        user.Account,
		Email:          user.Email,
		Firstname:      user.Firstname,
		Lastname:       user.Lastname,
		Status:         user.Status,
		Addr1:          user.Addr1,
		Addr2:          user.Addr2,
		City:           user.City,
		State:          user.State,
		Zip:            user.Zip,
		Country:        user.Country,
		LanguageCode:   user.LanguageCode,
		LastUpdateDate: time.Now(),
	}
	isSuccess := uc.UserService.UpdateUser(newUser)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_UPDATE),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SUCCESS_UPDATE),
		},
	}
}
