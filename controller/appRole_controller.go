package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

/**
 * 角色控制器结构体：用来实现处理角色模块的接口的请求，并返回给客户端
 */
type AppRoleController struct {
	//上下文对象
	Ctx iris.Context

	AppRoleService service.AppRoleService
	//session对象
	Session *sessions.Session
}

func (arc *AppRoleController) BeforeActivation(a mvc.BeforeActivation) {

	//添加角色
	a.Handle("POST", "/", "PostAddAppRole")

	//删除角色
	a.Handle("DELETE", "/{id}", "DeleteAppRole")

	//查询角色
	a.Handle("GET", "single/{id}", "GetAppRole")

	//查询角色数量
	a.Handle("GET", "/count", "GetCount")

	//所有角色列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询角色分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改角色
	a.Handle("PUT", "/", "UpdateAppRole")
}

/**
 * url: /appRole/addappRole
 * type：post
 * descs：添加角色
 */
func (arc *AppRoleController) PostAddAppRole() mvc.Result {

	var appRole model.TplAppRoleT
	err := arc.Ctx.ReadJSON(&appRole)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	newAppRole := &model.TplAppRoleT{
		Name:      appRole.Name,
		Desp:      appRole.Desp,
		Status:    appRole.Status,
		Defaultin: appRole.Defaultin,
		Code:      appRole.Code,
	}
	isSuccess := arc.AppRoleService.AddAppRole(newAppRole)
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
 * 删除角色
 */
func (arc *AppRoleController) DeleteAppRole() mvc.Result {

	id := arc.Ctx.Params().Get("id")

	appRoleId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := arc.AppRoleService.DeleteAppRole(appRoleId)
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
 * 获取角色
 * 请求类型：Get
 */
func (arc *AppRoleController) GetAppRole() mvc.Result {

	id := arc.Ctx.Params().Get("id")

	appRoleId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	appRole, err := arc.AppRoleService.GetAppRole(appRoleId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回角色
	return mvc.Response{
		Object: appRole.AppRoleTToRespDesc(),
	}
}

/**
 * 获取角色数量
 * 请求类型：Get
 */
func (arc *AppRoleController) GetCount() mvc.Result {

	//角色总数
	total, err := arc.AppRoleService.GetAppRoleTotalCount()

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
 * 获取角色列表
 * 请求类型：Get
 */
func (arc *AppRoleController) GetList() mvc.Result {

	appRoleList := arc.AppRoleService.GetAppRoleList()
	if len(appRoleList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的角色数据进行转换成前端需要的内容
	var respList []interface{}
	for _, appRole := range appRoleList {
		respList = append(respList, appRole.AppRoleTToRespDesc())
	}

	//返回角色列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取角色带参数分页查询
 * 请求类型：Get
 */
func (arc *AppRoleController) GetPageList() mvc.Result {

	offsetStr := arc.Ctx.FormValue("current")
	limitStr := arc.Ctx.FormValue("pageSize")
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
	} else {
		offset--
	}

	//做最大的限制
	if limit > MaxLimit {
		limit = MaxLimit
	}

	//获取页面参数
	name := arc.Ctx.FormValue("pId")
	appRoleParam := &model.TplAppRoleT{
		Name: name,
	}
	appRoleList := arc.AppRoleService.GetAppRolePageList(appRoleParam, offset, limit)
	total, _ := arc.AppRoleService.GetAppRoleTotalCount()
	if len(appRoleList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的角色数据进行转换成前端需要的内容
	var respList []interface{}
	for _, appRole := range appRoleList {
		respList = append(respList, appRole.AppRoleTToRespDesc())
	}

	//返回角色列表
	return mvc.Response{
		Object: map[string]interface{}{
			"data":     respList,
			"total":    total,
			"success":  true,
			"pageSize": limit,
			"current":  offset,
		},
	}
}

/**
 * type：put
 * descs：修改角色
 */
func (arc *AppRoleController) UpdateAppRole() mvc.Result {

	var appRole model.TplAppRoleT
	err := arc.Ctx.ReadJSON(&appRole)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}
	newAppRole := &model.TplAppRoleT{
		Id:        appRole.Id,
		Name:      appRole.Name,
		Desp:      appRole.Desp,
		Status:    appRole.Status,
		Defaultin: appRole.Defaultin,
		Code:      appRole.Code,
	}
	isSuccess := arc.AppRoleService.UpdateAppRole(newAppRole)
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
